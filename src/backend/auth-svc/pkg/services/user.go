package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
	auth "github.com/EdoRguez/business-manager/auth-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/auth-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/password_hash"
)

type UserService struct {
	Repo *repo.UserRepo
	auth.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *auth.CreateUserRequest) (*auth.CreateUserResponse, error) {
	fmt.Println("Auth Service :  CreateUser")
	fmt.Println("Auth Service :  CreateUser - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createUserParams := db.CreateUserParams{
		CompanyID:    req.ComanyId,
		RoleID:       req.RoleId,
		Email:        req.Email,
		PasswordHash: password_hash.HashPassword(req.Password),
	}

	c, err := s.Repo.CreateUser(ctx, createUserParams)
	if err != nil {
		fmt.Println("Auth Service :  CreateUser - ERROR")
		fmt.Println(err.Error())
		return &auth.CreateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  CreateUser - SUCCESS")
	return &auth.CreateUserResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *auth.GetUserRequest) (*auth.GetUserResponse, error) {
	fmt.Println("Auth Service :  GetUser")
	fmt.Println("Auth Service :  GetUser - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetUser(ctx, req.Id)
	if err != nil {
		fmt.Println("Auth Service :  GetUser - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &auth.GetUserResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  GetUser - SUCCESS")
	return &auth.GetUserResponse{
		Id:       c.ID,
		ComanyId: c.CompanyID,
		RoleId:   c.RoleID,
		Email:    c.Email,
		Status:   http.StatusOK,
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *auth.GetUsersRequest) (*auth.GetUsersResponse, error) {
	fmt.Println("Auth Service :  GetUsers")
	fmt.Println("Auth Service :  GetUsers - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetUsersParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	c, err := s.Repo.GetUsers(ctx, params)
	if err != nil {
		fmt.Println("Auth Service :  GetUsers - ERROR")
		fmt.Println(err.Error())
		return &auth.GetUsersResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	users := make([]*auth.GetUserResponse, 0, len(c))
	for _, v := range c {
		users = append(users, &auth.GetUserResponse{
			Id:       v.ID,
			ComanyId: v.CompanyID,
			RoleId:   v.RoleID,
			Email:    v.Email,
			Status:   http.StatusOK,
		})
	}

	fmt.Println("Auth Service :  GetUsers - SUCCESS")
	return &auth.GetUsersResponse{
		Users:  users,
		Status: http.StatusOK,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *auth.UpdateUserRequest) (*auth.UpdateUserResponse, error) {
	fmt.Println("Auth Service :  UpdateUser")
	fmt.Println("Auth Service :  UpdateUser - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdateUserParams{
		ID:           req.Id,
		RoleID:       req.RoleId,
		Email:        req.Email,
		PasswordHash: password_hash.HashPassword(req.Password),
	}

	_, err := s.Repo.UpdateUser(ctx, params)
	if err != nil {
		fmt.Println("Auth Service :  UpdateUser - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &auth.UpdateUserResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  UpdateUser - SUCCESS")
	return &auth.UpdateUserResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *auth.DeleteUserRequest) (*auth.DeleteUserResponse, error) {
	fmt.Println("Auth Service :  DeleteUser")
	fmt.Println("Auth Service :  DeleteUser - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeleteUser(ctx, req.Id)
	if err != nil {
		fmt.Println("Auth Service :  DeleteUser - ERROR")
		fmt.Println(err.Error())
		return &auth.DeleteUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  DeleteUser - SUCCESS")
	return &auth.DeleteUserResponse{
		Status: http.StatusNoContent,
	}, nil
}
