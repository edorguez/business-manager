package main

import (
	"fmt"
	"log"
	"net"

	"github.com/EdoRguez/business-manager/file-svc/pkg/config"
	pbfile "github.com/EdoRguez/business-manager/file-svc/pkg/pb/file"
	"github.com/EdoRguez/business-manager/file-svc/pkg/services"
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

	fmt.Println("Client Service ON: ", c.Port)

	ps := services.FileService{
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pbfile.RegisterFileServiceServer(grpcServer, &ps)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
