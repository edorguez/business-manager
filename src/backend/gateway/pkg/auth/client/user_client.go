package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/auth/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/auth/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"google.golang.org/grpc"
)

var userServiceClient pb.UserServiceClient

func InitUserServiceClient(c *config.Config) error {
	fmt.Println("User CLIENT :  InitUserServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Auth_Svc_Url, grpc.WithInsecure())

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

	fmt.Println("API Gateway :  CreateUser - SUCCESS")
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
