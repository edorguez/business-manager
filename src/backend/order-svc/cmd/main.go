package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", c.Port)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	mongoClient, err := db.ConnectMongoDb(c.DBSource)
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

	pb.RegisterProductServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
