package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"

	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/config"
	db "github.com/edorguez/business-manager/services/whatsapp-svc/pkg/db/sqlc"
	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/repository"
	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/services"
	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/wsmanager"
	pbwhatsapp "github.com/edorguez/business-manager/shared/pb/whatsapp"
	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// Create separate listeners for gRPC and HTTP
	grpcLis, err := net.Listen("tcp", ":"+c.WhatsappSvcPort)
	if err != nil {
		log.Fatalln("Failed to listen for gRPC:", err)
	}

	httpLis, err := net.Listen("tcp", ":"+c.WhatsappWsPort)
	if err != nil {
		log.Fatalln("Failed to listen for HTTP:", err)
	}

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development"
	}

	var dbSource string
	if appEnv == "development" {
		fmt.Println("Running in development mode")
		dbSource = c.WhatsappDBSourceDevelopment
	} else {
		fmt.Println("Running in docker mode")
		dbSource = c.WhatsappDBSourceDockerContainer
	}

	conn, err := sql.Open(c.PostgresDBDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	storage := db.NewStorage(conn)

	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)
	defer cancel()

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	container, err := sqlstore.New(ctx, c.PostgresDBDriver, dbSource, dbLog)
	if err != nil {
		panic(err)
	}

	setupAPI(ctx, container)

	fmt.Printf("gRPC Service ON: %s\n", c.WhatsappSvcPort)
	fmt.Printf("HTTP WebSocket Service ON: %s\n", c.WhatsappWsPort)

	ps := services.WhatsappService{
		Repo:   repository.NewWhatsappRepo(storage),
		Config: &c,
	}

	grpcServer := grpc.NewServer()
	pbwhatsapp.RegisterWhatsappServiceServer(grpcServer, &ps)

	// Run both servers concurrently
	go func() {
		log.Printf("Starting gRPC server on port %s", c.WhatsappSvcPort)
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve gRPC: %v", err)
		}
	}()

	go func() {
		log.Printf("Starting HTTP/WebSocket server on port %s", c.WhatsappWsPort)
		if err := http.Serve(httpLis, nil); err != nil {
			log.Fatalf("Failed to serve HTTP: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	fmt.Println("Shutting down servers...")
	grpcServer.GracefulStop()
	cancel()
}

// setupAPI will start all Routes and their Handlers
func setupAPI(ctx context.Context, container *sqlstore.Container) {

	// Create a Manager instance used to handle WebSocket Connections
	manager := wsmanager.NewManager(ctx, *container)

	// Serve the ./frontend directory at Route /
	// http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/login", manager.LoginHandler)
	http.HandleFunc("/ws", manager.ServeWS)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.Clients))
	})
}
