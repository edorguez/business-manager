package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"

	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
	pb "github.com/EdoRguez/business-manager/auth-svc/pkg/pb/user"
	repo "github.com/EdoRguez/business-manager/auth-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/password_hash"
)

type UserService struct {
	Repo *repo.UserRepo
	pb.UnimplementedUserServiceServer
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	fmt.Println("Auth Service :  CreateUser")
	fmt.Println("Auth Service :  CreateUser - Req")
	// fmt.Println(req)
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  CreateUser - ERROR")
		fmt.Println(err.Error())
		return &pb.CreateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if u.Email == req.Email {
		fmt.Println("Auth Service :  CreateUser - ERROR")
		fmt.Println("User already exists")
		return &pb.CreateUserResponse{
			Status: http.StatusInternalServerError,
			Error:  "User already exists",
		}, nil
	}

	createUserParams := db.CreateUserParams{
		CompanyID:    req.CompanyId,
		RoleID:       req.RoleId,
		Email:        req.Email,
		PasswordHash: password_hash.HashPassword(req.Password),
	}

	c, err := s.Repo.CreateUser(ctx, createUserParams)
	if err != nil {
		fmt.Println("Auth Service :  CreateUser - ERROR")
		fmt.Println(err.Error())
		return &pb.CreateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  CreateUser - SUCCESS")
	return &pb.CreateUserResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
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

		return &pb.GetUserResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  GetUser - SUCCESS")
	return &pb.GetUserResponse{
		Id:        c.ID,
		CompanyId: c.CompanyID,
		RoleId:    c.RoleID,
		Email:     c.Email,
		Role: &pb.Role{
			Id:   c.AuthRole.ID,
			Name: c.AuthRole.Name,
		},
		Status: http.StatusOK,
	}, nil
}

func (s *UserService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	fmt.Println("Auth Service :  GetUsers")
	fmt.Println("Auth Service :  GetUsers - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetUsersParams{
		CompanyID: req.CompanyId,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	c, err := s.Repo.GetUsers(ctx, params)
	if err != nil {
		fmt.Println("Auth Service :  GetUsers - ERROR")
		fmt.Println(err.Error())
		return &pb.GetUsersResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	users := make([]*pb.GetUserResponse, 0, len(c))
	for _, v := range c {
		users = append(users, &pb.GetUserResponse{
			Id:        v.ID,
			CompanyId: v.CompanyID,
			RoleId:    v.RoleID,
			Email:     v.Email,
			Role: &pb.Role{
				Id:   v.AuthRole.ID,
				Name: v.AuthRole.Name,
			},
			Status: http.StatusOK,
		})
	}

	fmt.Println("Auth Service :  GetUsers - SUCCESS")
	return &pb.GetUsersResponse{
		Users:  users,
		Status: http.StatusOK,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	fmt.Println("Auth Service :  UpdateUser")
	fmt.Println("Auth Service :  UpdateUser - Req")
	// fmt.Println(req)
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  Updateuser - ERROR")
		fmt.Println(err.Error())
		return &pb.UpdateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if u.Email == req.Email && u.ID != req.Id {
		fmt.Println("Auth Service :  UpdateUser - ERROR")
		fmt.Println("User already exists")
		return &pb.UpdateUserResponse{
			Status: http.StatusInternalServerError,
			Error:  "User already exists",
		}, nil
	}

	oldUser, errOldUser := s.Repo.GetUser(ctx, req.Id)
	if errOldUser != nil {
		fmt.Println("Auth Service :  Updateuser - ERROR")
		fmt.Println(err.Error())
		return &pb.UpdateUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	// req.Password is an optional parameter, if we don't send it, then, set the original value
	passwordHash := oldUser.PasswordHash
	if req.Password != nil {
		passwordHash = password_hash.HashPassword(*req.Password)
	}

	params := db.UpdateUserParams{
		ID:           req.Id,
		RoleID:       req.RoleId,
		Email:        req.Email,
		PasswordHash: passwordHash,
	}

	_, err = s.Repo.UpdateUser(ctx, params)
	if err != nil {
		fmt.Println("Auth Service :  UpdateUser - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &pb.UpdateUserResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  UpdateUser - SUCCESS")
	return &pb.UpdateUserResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
	fmt.Println("Auth Service :  DeleteUser")
	fmt.Println("Auth Service :  DeleteUser - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeleteUser(ctx, req.Id)
	if err != nil {
		fmt.Println("Auth Service :  DeleteUser - ERROR")
		fmt.Println(err.Error())
		return &pb.DeleteUserResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  DeleteUser - SUCCESS")
	return &pb.DeleteUserResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *UserService) UpdateEmail(ctx context.Context, req *pb.UpdateEmailRequest) (*pb.UpdateEmailResponse, error) {
	fmt.Println("Auth Service :  UpdateEmail")
	fmt.Println("Auth Service :  UpdateEmail - Req")
	// fmt.Println(req)
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  UpdateEmail - ERROR")
		fmt.Println(err.Error())
		return &pb.UpdateEmailResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if strings.ToLower(u.Email) == strings.ToLower(req.Email) {
		fmt.Println("Auth Service :  UpdateEmail - ERROR")
		fmt.Println("Email already exists")
		return &pb.UpdateEmailResponse{
			Status: http.StatusInternalServerError,
			Error:  "Email already exists",
		}, nil
	}

	params := db.UpdateEmailParams{
		ID:    req.Id,
		Email: req.Email,
	}

	_, err = s.Repo.UpdateEmail(ctx, params)
	if err != nil {
		fmt.Println("Auth Service :  UpdateEmail - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &pb.UpdateEmailResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  UpdateEmail - SUCCESS")
	return &pb.UpdateEmailResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordRequest) (*pb.UpdatePasswordResponse, error) {
	fmt.Println("Auth Service :  UpdatePassword")
	fmt.Println("Auth Service :  UpdatePassword - Req")
	// fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdatePasswordParams{
		ID:           req.Id,
		PasswordHash: password_hash.HashPassword(req.Password),
	}

	_, err := s.Repo.UpdatePassword(ctx, params)
	if err != nil {
		fmt.Println("Auth Service :  UpdatePassword - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &pb.UpdatePasswordResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Auth Service :  UpdatePassword - SUCCESS")
	return &pb.UpdatePasswordResponse{
		Status: http.StatusNoContent,
	}, nil
}
