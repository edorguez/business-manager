package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/EdoRguez/business-manager/auth-svc/pkg/config"
	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
	pbauth "github.com/EdoRguez/business-manager/auth-svc/pkg/pb/auth"
	pbrole "github.com/EdoRguez/business-manager/auth-svc/pkg/pb/role"
	pbuser "github.com/EdoRguez/business-manager/auth-svc/pkg/pb/user"
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

	as := services.AuthService{
		Repo: repo.NewUserRepo(storage),
		Jwt: jwt_manager.JWTWrapper{
			SecretKey:       c.JWTSecretKey,
			Issuer:          "auth-svc",
			ExpirationHours: 24 * 7,
		},
		Config: &c,
	}

	us := services.UserService{
		Repo: repo.NewUserRepo(storage),
	}

	ro := services.RoleService{
		Repo: repo.NewRoleRepo(storage),
	}

	grpcServer := grpc.NewServer()

	pbauth.RegisterAuthServiceServer(grpcServer, &as)
	pbuser.RegisterUserServiceServer(grpcServer, &us)
	pbrole.RegisterRoleServiceServer(grpcServer, &ro)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
