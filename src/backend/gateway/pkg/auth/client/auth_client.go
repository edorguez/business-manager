package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/auth"
	"google.golang.org/grpc"
)

var authServiceClient pb.AuthServiceClient

func InitAuthServiceClient(c *config.Config) error {
	fmt.Println("Auth CLIENT :  InitAuthServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var authSvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		authSvcUrl = c.Auth_Svc_Url + ":" + c.Auth_Svc_Port
	} else {
		fmt.Println("Running in development mode")
		authSvcUrl = c.Development_Url + ":" + c.Auth_Svc_Port
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(authSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	authServiceClient = pb.NewAuthServiceClient(cc)
	return nil
}

func Register(body contracts.CreateUserRequest, c context.Context) (*pb.RegisterResponse, *contracts.Error) {
	fmt.Println("Auth CLIENT :  Register")

	fmt.Println("Auth CLIENT :  Register - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	registerParams := &pb.RegisterRequest{
		CompanyId: body.CompanyId,
		RoleId:    body.RoleId,
		Email:     body.Email,
		Password:  body.Password,
	}

	res, err := authServiceClient.Register(c, registerParams)

	if err != nil {
		fmt.Println("Auth CLIENT :  Register - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Auth CLIENT :  Register - SUCCESS")
	return res, nil
}

func Login(body contracts.LoginRequest, c context.Context) (*pb.LoginResponse, *contracts.Error) {
	fmt.Println("Auth CLIENT :  Login")

	fmt.Println("Auth CLIENT :  Login - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	loginParams := &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := authServiceClient.Login(c, loginParams)

	if err != nil {
		fmt.Println("Auth CLIENT :  Login - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Auth CLIENT :  Login - SUCCESS")
	return res, nil
}

func Validate(body contracts.ValidateRequest, c context.Context) (*pb.ValidateResponse, *contracts.Error) {
	fmt.Println("Auth CLIENT :  Login")

	fmt.Println("Auth CLIENT :  Login - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	validateParams := &pb.ValidateRequest{
		Token: body.Token,
	}

	res, err := authServiceClient.Validate(c, validateParams)

	if err != nil {
		fmt.Println("Auth CLIENT :  Validate - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Auth CLIENT :  Validate - SUCCESS")
	return res, nil
}
