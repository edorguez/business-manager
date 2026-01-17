package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/edorguez/business-manager/services/order-svc/pkg/config"
	db "github.com/edorguez/business-manager/services/order-svc/pkg/db/sqlc"
	"github.com/edorguez/business-manager/services/order-svc/pkg/repository"
	"github.com/edorguez/business-manager/services/order-svc/pkg/services"
	pborder "github.com/edorguez/business-manager/shared/pb/order"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", ":"+c.OrderSvcPort)
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
		dbSource = c.OrderDBSourceDevelopment
	} else {
		fmt.Println("Running in docker mode")
		dbSource = c.OrderDBSourceDockerContainer
	}

	conn, err := sql.Open(c.PostgresDBDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	storage := db.NewStorage(conn)

	fmt.Println("Order Service ON: ", c.OrderSvcPort)

	ps := services.OrderService{
		Repo:   repository.NewOrderRepo(storage),
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pborder.RegisterOrderServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
