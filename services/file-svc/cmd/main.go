package main

import (
	"fmt"
	"log"
	"net"

	"github.com/edorguez/business-manager/services/file-svc/pkg/config"
	"github.com/edorguez/business-manager/services/file-svc/pkg/services"
	pbfile "github.com/edorguez/business-manager/shared/pb/file"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	lis, err := net.Listen("tcp", ":"+c.FileSvcPort)
	if err != nil {
		log.Fatalln("Failed to listing:", err)
	}

	fmt.Println("Client Service ON: ", c.FileSvcPort)

	s := services.FileService{
		Config: &c,
	}

	grpcServer := grpc.NewServer()

	pbfile.RegisterFileServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
