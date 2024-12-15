package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	"github.com/EdoRguez/business-manager/order-svc/pkg/pb"
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

	pb.RegisterOrderServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
