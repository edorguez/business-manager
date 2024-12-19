package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/user"
	"google.golang.org/grpc"
)

var userServiceClient pb.UserServiceClient

func InitUserServiceClient(c *config.Config) error {
	fmt.Println("User CLIENT :  InitUserServiceClient")

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

	userServiceClient = pb.NewUserServiceClient(cc)
	return nil
}

func CreateUser(body contracts.CreateUserRequest, c context.Context) (*pb.CreateUserResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  CreateUser")

	fmt.Println("User CLIENT :  CreateUser - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createUserParams := &pb.CreateUserRequest{
		CompanyId: body.CompanyId,
		RoleId:    body.RoleId,
		Email:     body.Email,
		Password:  body.Password,
	}

	res, err := userServiceClient.CreateUser(c, createUserParams)

	if err != nil {
		fmt.Println("User CLIENT :  CreateUser - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("User CLIENT :  CreateUser - SUCCESS")
	return res, nil
}

func GetUser(id int64, c context.Context) (*contracts.GetUserResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  GetUser")

	params := &pb.GetUserRequest{
		Id: id,
	}

	res, err := userServiceClient.GetUser(c, params)

	if err != nil {
		fmt.Println("User CLIENT :  GetUser - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("User CLIENT :  GetUser - SUCCESS")

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetUserResponse{
		Id:        res.Id,
		CompanyId: res.CompanyId,
		RoleId:    res.RoleId,
		Email:     res.Email,
	}, nil
}

func GetUsers(params *pb.GetUsersRequest, c context.Context) ([]*contracts.GetUserResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  GetUser")

	res, err := userServiceClient.GetUsers(c, params)

	if err != nil {
		fmt.Println("User CLIENT :  GetUsers - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	cr := make([]*contracts.GetUserResponse, 0, len(res.Users))
	for _, v := range res.Users {
		cr = append(cr, &contracts.GetUserResponse{
			Id:        v.Id,
			CompanyId: v.CompanyId,
			RoleId:    v.RoleId,
			Email:     v.Email,
			Role: contracts.GetRoleResponse{
				Id:   v.Role.Id,
				Name: v.Role.Name,
			},
		})
	}

	fmt.Println("User CLIENT :  GetUsers - SUCCESS")
	return cr, nil
}

func UpdateUser(id int64, body contracts.UpdateUserRequest, c context.Context) (*pb.UpdateUserResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  UpdateUser")

	fmt.Println("User CLIENT :  UpdateUser - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateUserParams := &pb.UpdateUserRequest{
		Id:       id,
		RoleId:   body.RoleId,
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := userServiceClient.UpdateUser(c, updateUserParams)

	if err != nil {
		fmt.Println("User CLIENT :  UpdateUser - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("User CLIENT :  UpdateUser - SUCCESS")
	return res, nil
}

func DeleteUser(id int64, c context.Context) (*pb.DeleteUserResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  DeleteUser")

	params := &pb.DeleteUserRequest{
		Id: id,
	}

	res, err := userServiceClient.DeleteUser(c, params)

	if err != nil {
		fmt.Println("User CLIENT :  DeleteUser - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("User CLIENT :  DeleteUser - SUCCESS")
	return res, nil
}

func UpdateEmail(id int64, body contracts.UpdateEmailRequest, c context.Context) (*pb.UpdateEmailResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  UpdateEmail")

	fmt.Println("User CLIENT :  UpdateEmail - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateUserParams := &pb.UpdateEmailRequest{
		Id:    id,
		Email: body.Email,
	}

	res, err := userServiceClient.UpdateEmail(c, updateUserParams)

	if err != nil {
		fmt.Println("User CLIENT :  UpdateEmail - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("User CLIENT :  UpdateEmail - SUCCESS")
	return res, nil
}

func UpdatePassword(id int64, body contracts.UpdatePasswordRequest, c context.Context) (*pb.UpdatePasswordResponse, *contracts.Error) {
	fmt.Println("User CLIENT :  UpdatePassword")

	fmt.Println("User CLIENT :  UpdatePassword - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateUserParams := &pb.UpdatePasswordRequest{
		Id:       id,
		Password: body.Password,
	}

	res, err := userServiceClient.UpdatePassword(c, updateUserParams)

	if err != nil {
		fmt.Println("User CLIENT :  UpdatePassword - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("User CLIENT :  UpdatePassword - SUCCESS")
	return res, nil
}
