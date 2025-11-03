package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/gateway/pkg/auth/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/auth"
	"github.com/edorguez/business-manager/shared/types"
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
		authSvcUrl = c.AuthSvcUrl + ":" + c.AuthSvcPort
	} else {
		fmt.Println("Running in development mode")
		authSvcUrl = c.DevelopmentUrl + ":" + c.AuthSvcPort
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

func SignUp(body contracts.SignUpRequest, images [][]byte, c context.Context) (*pb.SignUpResponse, *types.Error) {
	fmt.Println("Auth CLIENT :  Sign Up")

	fmt.Println("Auth CLIENT :  Sign Up - Body")
	// fmt.Println(body)
	fmt.Println("-----------------")

	signUpParams := &pb.SignUpRequest{
		Company: &pb.SignUpCompany{
			Name:   body.Company.Name,
			Phone:  body.Company.Phone,
			Images: images,
		},
		User: &pb.SignUpUser{
			Email:    body.User.Email,
			Password: body.User.Password,
		},
	}

	res, err := authServiceClient.SignUp(c, signUpParams)

	if err != nil {
		fmt.Println("Auth CLIENT :  Sign Up - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Auth CLIENT :  Sign Up - SUCCESS")
	return res, nil
}

func Login(body contracts.LoginRequest, c context.Context) (*pb.LoginResponse, *types.Error) {
	fmt.Println("Auth CLIENT :  Login")

	fmt.Println("Auth CLIENT :  Login - Body")
	// fmt.Println(body)
	fmt.Println("-----------------")

	loginParams := &pb.LoginRequest{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := authServiceClient.Login(c, loginParams)

	if err != nil {
		fmt.Println("Auth CLIENT :  Login - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Auth CLIENT :  Login - SUCCESS")
	return res, nil
}

func Validate(body contracts.ValidateRequest, c context.Context) (*pb.ValidateResponse, *types.Error) {
	fmt.Println("Auth CLIENT :  Login")

	fmt.Println("Auth CLIENT :  Login - Body")
	// fmt.Println(body)
	fmt.Println("-----------------")

	validateParams := &pb.ValidateRequest{
		Token: body.Token,
	}

	res, err := authServiceClient.Validate(c, validateParams)

	if err != nil {
		fmt.Println("Auth CLIENT :  Validate - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Auth CLIENT :  Validate - SUCCESS")
	return res, nil
}
