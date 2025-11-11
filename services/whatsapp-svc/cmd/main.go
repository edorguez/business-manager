package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/config"
	db "github.com/edorguez/business-manager/services/whatsapp-svc/pkg/db/sqlc"
	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/repository"
	"github.com/edorguez/business-manager/services/whatsapp-svc/pkg/services"
	pbwhatsapp "github.com/edorguez/business-manager/shared/pb/whatsapp"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", ":"+c.WhatsappSvcPort)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
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

	fmt.Println("Client Service ON: ", c.WhatsappSvcPort)

	ps := services.WhatsappService{
		Repo:   repository.NewWhatsappRepo(storage),
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pbwhatsapp.RegisterWhatsappServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}

// func main() {
// 	c, err := config.LoadConfig()
// 	if err != nil {
// 		log.Fatalln("Failed at config", err)
// 	}

// 	dbLog := waLog.Stdout("Database", "DEBUG", true)
// 	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
// 	container, err := sqlstore.New(c.DBDriver, c.DBSource, dbLog)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Create a root ctx and a CancelFunc which can be used to cancel retentionMap goroutine
// 	rootCtx := context.Background()
// 	ctx, cancel := context.WithCancel(rootCtx)

// 	defer cancel()

// 	setupAPI(ctx, container)

// 	// Serve on port :50055, fudge yeah hardcoded port
// 	// err := http.ListenAndServeTLS(":50055", "server.crt", "server.key", nil)
// 	err = http.ListenAndServe(":"+c.Port, nil)
// 	if err != nil {
// 		log.Fatal("ListenAndServe: ", err)
// 	}
// }

// // setupAPI will start all Routes and their Handlers
// func setupAPI(ctx context.Context, container *sqlstore.Container) {

// 	// Create a Manager instance used to handle WebSocket Connections
// 	manager := wsmanager.NewManager(ctx, *container)

// 	// Serve the ./frontend directory at Route /
// 	// http.Handle("/", http.FileServer(http.Dir("./frontend")))
// 	http.HandleFunc("/login", manager.LoginHandler)
// 	http.HandleFunc("/ws", manager.ServeWS)

// 	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, len(manager.Clients))
// 	})
// }
