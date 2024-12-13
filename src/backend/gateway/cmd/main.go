package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth"
	"github.com/EdoRguez/business-manager/gateway/pkg/company"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer"
	"github.com/EdoRguez/business-manager/gateway/pkg/order"
	"github.com/EdoRguez/business-manager/gateway/pkg/product"
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

	// run the server
	err = startServer(conf.Gateway_Url, &conf)
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

	s := &http.Server{
		Addr:         address,
		Handler:      rc,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
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
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

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
