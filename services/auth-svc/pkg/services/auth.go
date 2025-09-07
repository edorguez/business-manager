package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/EdoRguez/business-manager/auth-svc/pkg/client"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/config"
	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
	pb "github.com/EdoRguez/business-manager/auth-svc/pkg/pb/auth"
	repo "github.com/EdoRguez/business-manager/auth-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/jwt_manager"
	"github.com/EdoRguez/business-manager/auth-svc/pkg/util/password_hash"
)

type AuthService struct {
	Repo   *repo.UserRepo
	Jwt    jwt_manager.JWTWrapper
	Config *config.Config
	pb.UnimplementedAuthServiceServer
}

func (s *AuthService) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	fmt.Println("Auth Service :  Register")
	fmt.Println("Auth Service :  Register - Req")
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  Register - ERROR")
		fmt.Println(err.Error())
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if u.Email == req.Email {
		fmt.Println("Auth Service :  Register - ERROR")
		fmt.Println("User already exists")
		return &pb.RegisterResponse{
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

	_, err = s.Repo.CreateUser(ctx, createUserParams)
	if err != nil {
		fmt.Println("Auth Service :  Register - ERROR")
		fmt.Println(err.Error())
		return &pb.RegisterResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  Register - SUCCESS")
	return &pb.RegisterResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *AuthService) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	fmt.Println("Auth Service :  Login")
	fmt.Println("Auth Service :  Login - Req")
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  Login - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Email or password incorrect"
		}

		return &pb.LoginResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	match := password_hash.CheckPasswordHash(req.Password, u.PasswordHash)

	if !match {
		return &pb.LoginResponse{
			Status: http.StatusNotFound,
			Error:  "Email or password incorrect",
		}, nil
	}

	if err := client.InitCompanyServiceClient(s.Config); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	company, errCompany := client.GetCompany(u.CompanyID, ctx)

	if errCompany != nil {
		fmt.Println("Auth Service :  Login - ERROR")
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  errCompany.Error(),
		}, nil
	}

	if company.LastPaymentDate.AsTime().Before(time.Now()) {
		fmt.Println("Auth Service :  Login - ERROR")
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Can't log in because of last payment, please contact support",
		}, nil
	}

	token, err := s.Jwt.GenerateToken(u.ID, u.Email, u.RoleID, u.CompanyID, company.PlanId)
	if err != nil {
		fmt.Println("Auth Service :  Login - ERROR")
		fmt.Println(err.Error())
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  Login - SUCCESS")
	return &pb.LoginResponse{
		Status: http.StatusOK,
		Token:  token,
	}, nil
}

func (s *AuthService) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	fmt.Println("Auth Service :  Validate")
	fmt.Println("Auth Service :  Validate - Req")
	fmt.Println("----------------")

	err := s.Jwt.ValidateToken(req.Token)

	if err != nil {
		return &pb.ValidateResponse{
			Status: http.StatusBadRequest,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  Validate - SUCCESS")
	return &pb.ValidateResponse{
		Status: http.StatusOK,
	}, nil
}
