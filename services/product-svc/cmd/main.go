package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/EdoRguez/business-manager/product-svc/pkg/config"
	"github.com/EdoRguez/business-manager/product-svc/pkg/db"
	pbproduct "github.com/EdoRguez/business-manager/product-svc/pkg/pb/product"
	repo "github.com/EdoRguez/business-manager/product-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/product-svc/pkg/services"
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

	mongoClient, err := db.ConnectMongoDb(dbSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer mongoClient.Disconnect(context.Background())

	fmt.Println("Client Service ON: ", c.Port)

	ps := services.ProductService{
		Repo:   repo.NewProductRepo(mongoClient, c),
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pbproduct.RegisterProductServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
