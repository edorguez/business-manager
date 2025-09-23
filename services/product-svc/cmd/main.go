package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/edorguez/business-manager/services/product-svc/pkg/config"
	"github.com/edorguez/business-manager/services/product-svc/pkg/db"
	repo "github.com/edorguez/business-manager/services/product-svc/pkg/repository"
	"github.com/edorguez/business-manager/services/product-svc/pkg/services"
	pbproduct "github.com/edorguez/business-manager/shared/pb/product"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", ":"+c.ProductSvcPort)
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
		dbSource = c.ProductDBSourceProduction
	} else {
		fmt.Println("Running in development mode")
		dbSource = c.ProductDBSourceDevelopment
	}

	mongoClient, err := db.ConnectMongoDb(dbSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer mongoClient.Disconnect(context.Background())

	fmt.Println("Client Service ON: ", c.ProductSvcPort)

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
