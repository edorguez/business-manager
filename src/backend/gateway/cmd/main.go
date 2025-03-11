package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth"
	"github.com/EdoRguez/business-manager/gateway/pkg/company"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer"
	"github.com/EdoRguez/business-manager/gateway/pkg/order"
	"github.com/EdoRguez/business-manager/gateway/pkg/product"
	"github.com/EdoRguez/business-manager/gateway/pkg/whatsapp"
	"github.com/gorilla/mux"
)

func main() {
	var exitCode int
	defer func() {
		os.Exit(exitCode)
	}()

	// load config
	conf, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		exitCode = 1
		return
	}

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var gatewayUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		gatewayUrl = conf.Production_Url + ":" + conf.Gateway_Port
	} else {
		fmt.Println("Running in development mode")
		gatewayUrl = conf.Development_Url + ":" + conf.Gateway_Port
	}

	// run the server
	err = startServer(gatewayUrl, &conf)
	if err != nil {
		log.Fatal("Cannot start server: ", err)
	}
}

func startServer(address string, conf *config.Config) error {
	sm := mux.NewRouter()
	rc := routesConfig(sm)

	baseRoute := sm.PathPrefix("/api").Subrouter()

	customer.LoadRoutes(baseRoute, conf)
	company.LoadRoutes(baseRoute, conf)
	product.LoadRoutes(baseRoute, conf)
	auth.LoadRoutes(baseRoute, conf)
	order.LoadRoutes(baseRoute, conf)
	whatsapp.LoadRoutes(baseRoute, conf)

	s := &http.Server{
		Addr:         address,
		Handler:      rc,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		fmt.Println("Starting server ON: ", address)

		err := s.ListenAndServe()
		if err != nil {
			fmt.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(ctx)
	return nil
}

func routesConfig(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		fmt.Println("--> Received Origin:", origin)

		// Check if the origin is from your domain or its subdomains
		if isValidOrigin(origin) {
			// Set the Access-Control-Allow-Origin header to the specific origin
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		// Set JSON Content-Type
		w.Header().Set("Content-Type", "application/json")

		// Check if the request is for CORS preflight
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Pass down the request to the next middleware (or final handler)
		next.ServeHTTP(w, r)
	})
}

func isValidOrigin(origin string) bool {
	// Define your main domain
	mainDomain := "www.edezco.com"

	// Check if the origin is exactly your main domain
	if origin == "https://"+mainDomain || origin == "http://"+mainDomain {
		return true
	}

	// Check if the origin is a subdomain of your main domain
	if strings.HasSuffix(origin, "."+mainDomain) {
		return true
	}

	// If the origin doesn't match, return false
	return false
}
