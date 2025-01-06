package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/EdoRguez/business-manager/company-svc/pkg/config"
	db "github.com/EdoRguez/business-manager/company-svc/pkg/db/sqlc"
	pbcompany "github.com/EdoRguez/business-manager/company-svc/pkg/pb/company"
	pbpayment "github.com/EdoRguez/business-manager/company-svc/pkg/pb/payment"
	pbpaymenttype "github.com/EdoRguez/business-manager/company-svc/pkg/pb/payment_type"
	repo "github.com/EdoRguez/business-manager/company-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/company-svc/pkg/services"
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

	cs := services.CompanyService{
		Repo:   repo.NewCompanyRepo(storage),
		Config: &c,
	}

	ps := services.PaymentService{
		Repo:        repo.NewPaymentRepo(storage),
		CompanyRepo: repo.NewCompanyRepo(storage),
	}

	pts := services.PaymentTypeService{
		Repo: repo.NewPaymentTypeRepo(storage),
	}

	grpcServer := grpc.NewServer()

	pbcompany.RegisterCompanyServiceServer(grpcServer, &cs)
	pbpayment.RegisterPaymentServiceServer(grpcServer, &ps)
	pbpaymenttype.RegisterPaymentTypeServiceServer(grpcServer, &pts)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
