package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	pborder "github.com/EdoRguez/business-manager/order-svc/pkg/pb/order"
	"github.com/EdoRguez/business-manager/order-svc/pkg/services"
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

	fmt.Println("Client Service ON: ", c.Port)

	ps := services.OrderService{
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pborder.RegisterOrderServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
