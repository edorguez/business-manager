package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/edorguez/business-manager/services/customer-svc/pkg/config"
	db "github.com/edorguez/business-manager/services/customer-svc/pkg/db/sqlc"
	repo "github.com/edorguez/business-manager/services/customer-svc/pkg/repository"
	"github.com/edorguez/business-manager/services/customer-svc/pkg/services"
	pbcustomer "github.com/edorguez/business-manager/shared/pb/customer"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", ":"+c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var dbSource string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		dbSource = c.DBSourceProduction
	} else {
		fmt.Println("Running in development mode")
		dbSource = c.DBSourceDevelopment
	}

	conn, err := sql.Open(c.DBDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	storage := db.NewStorage(conn)

	fmt.Println("Client Service ON: ", c.Port)

	s := services.CustomerService{
		Repo: repo.NewCustomerRepo(storage),
	}

	grpcServer := grpc.NewServer()

	pbcustomer.RegisterCustomerServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
