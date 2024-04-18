package main

import (
	"context"
	"fmt"
	"log"

	"github.com/EdoRguez/business-manager/product-svc/pkg/config"
	"github.com/EdoRguez/business-manager/product-svc/pkg/db"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	// lis, err := net.Listen("tcp", c.Port)
	// if err != nil {
	// 	log.Fatalln("Failed to listing:", err)
	// }

	mongoClient, err := db.ConnectMongoDb(c.DBSource)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer mongoClient.Disconnect(context.Background())

	// storage := db.NewStorage(conn)

	fmt.Println("Client Service ON: ", c.Port)

	// cs := services.CompanyService{
	// 	Repo: repo.NewCompanyRepo(storage),
	// }
	//
	// ps := services.PaymentService{
	// 	Repo: repo.NewPaymentRepo(storage),
	// }
	//
	// grpcServer := grpc.NewServer()
	//
	// pb.RegisterCompanyServiceServer(grpcServer, &cs)
	// pb.RegisterPaymentServiceServer(grpcServer, &ps)

	// if err := grpcServer.Serve(lis); err != nil {
	// 	log.Fatalln("Failed to serve:", err)
	// }
}
