package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/EdoRguez/business-manager/auth-svc/pkg/config"
	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/auth-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/services"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/jwt_manager"
	_ "github.com/lib/pq"
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

	conn, err := sql.Open(c.DBDriver, c.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db: ", err)
	}

	storage := db.NewStorage(conn)

	fmt.Println("Client Service ON: ", c.Port)

	as := services.AuthService{
		Repo: repo.NewUserRepo(storage),
		Jwt: jwt_manager.JWTWrapper{
			SecretKey:       c.JWTSecretKey,
			Issuer:          "auth-svc",
			ExpirationHours: 24 * 365,
		},
	}

	us := services.UserService{
		Repo: repo.NewUserRepo(storage),
	}

	grpcServer := grpc.NewServer()

	pb.RegisterAuthServiceServer(grpcServer, &as)
	pb.RegisterUserServiceServer(grpcServer, &us)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
