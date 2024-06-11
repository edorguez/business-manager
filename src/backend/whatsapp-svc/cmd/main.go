package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/config"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/services"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/ws"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	dbLog := waLog.Stdout("Database", "DEBUG", true)
	// Make sure you add appropriate DB connector imports, e.g. github.com/mattn/go-sqlite3 for SQLite
	container, err := sqlstore.New(c.DBDriver, c.DBSource, dbLog)
	if err != nil {
		panic(err)
	}

	// Create a root ctx and a CancelFunc which can be used to cancel retentionMap goroutine
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()

	setupAPI(ctx, container)

	go services.StartWhatsapp(container)

	// Serve on port :50055, fudge yeah hardcoded port
	// err := http.ListenAndServeTLS(":50055", "server.crt", "server.key", nil)
	err = http.ListenAndServe(":"+c.Port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

// setupAPI will start all Routes and their Handlers
func setupAPI(ctx context.Context, container *sqlstore.Container) {

	// Create a Manager instance used to handle WebSocket Connections
	manager := ws.NewManager(ctx, container)

	// Serve the ./frontend directory at Route /
	// http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/login", manager.LoginHandler)
	http.HandleFunc("/ws", manager.ServeWS)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.Clients))
	})
}
