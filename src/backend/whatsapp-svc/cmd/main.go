package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg"
)

func main() {

	// Create a root ctx and a CancelFunc which can be used to cancel retentionMap goroutine
	rootCtx := context.Background()
	ctx, cancel := context.WithCancel(rootCtx)

	defer cancel()

	setupAPI(ctx)

	// Serve on port :8080, fudge yeah hardcoded port
	// err := http.ListenAndServeTLS(":8080", "server.crt", "server.key", nil)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

	fmt.Println("Started")
}

// setupAPI will start all Routes and their Handlers
func setupAPI(ctx context.Context) {

	// Create a Manager instance used to handle WebSocket Connections
	manager := pkg.NewManager(ctx)

	// Serve the ./frontend directory at Route /
	// http.Handle("/", http.FileServer(http.Dir("./frontend")))
	// http.HandleFunc("/login", manager.LoginHandler)
	http.HandleFunc("/ws", manager.ServeWS)

	http.HandleFunc("/debug", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, len(manager.Clients))
	})
}
