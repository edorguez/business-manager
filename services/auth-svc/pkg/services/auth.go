package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/edorguez/business-manager/services/auth-svc/pkg/client"
	"github.com/edorguez/business-manager/services/auth-svc/pkg/config"
	"github.com/edorguez/business-manager/services/auth-svc/pkg/constants"
	db "github.com/edorguez/business-manager/services/auth-svc/pkg/db/sqlc"
	repo "github.com/edorguez/business-manager/services/auth-svc/pkg/repository"
	pb "github.com/edorguez/business-manager/shared/pb/auth"
	pbcompany "github.com/edorguez/business-manager/shared/pb/company"
	"github.com/edorguez/business-manager/shared/util/jwt_manager"
	"github.com/edorguez/business-manager/shared/util/password_hash"
)

type AuthService struct {
	Repo   *repo.UserRepo
	Jwt    jwt_manager.JWTWrapper
	Config *config.Config
	pb.UnimplementedAuthServiceServer
}

func (s *AuthService) SignUp(ctx context.Context, req *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	fmt.Println("Auth Service :  Sign Up")
	fmt.Println("Auth Service :  Sign Up - Req")
	fmt.Println("----------------")

	u, err := s.Repo.GetUserByEmail(ctx, req.User.Email)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("Auth Service :  Register - ERROR")
		fmt.Println(err.Error())
		return &pb.SignUpResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	if u.Email == req.User.Email {
		fmt.Println("Auth Service :  Sign Up - ERROR")
		fmt.Println("User email already exists")
		return &pb.SignUpResponse{
			Status: http.StatusInternalServerError,
			Error:  "User email already exists",
		}, nil
	}

	if err := client.InitCompanyServiceClient(s.Config); err != nil {
		return &pb.SignUpResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	companyParams := &pbcompany.CreateCompanyRequest{
		Name: req.Company.Name,
	}

	company, errCompany := client.CreateCompany(companyParams, ctx)

	if errCompany != nil {
		fmt.Println("Auth Service :  Sign Up - ERROR")
		return &pb.SignUpResponse{
			Status: http.StatusInternalServerError,
			Error:  errCompany.Error(),
		}, nil
	}

	createUserParams := db.CreateUserParams{
		CompanyID:    company.Id,
		RoleID:       constants.ROLE_ID_ADMIN,
		Email:        req.User.Email,
		PasswordHash: password_hash.HashPassword(req.User.Password),
	}

	_, err = s.Repo.CreateUser(ctx, createUserParams)
	if err != nil {
		fmt.Println("Auth Service :  Sign Up - ERROR")
		fmt.Println(err.Error())
		return &pb.SignUpResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Auth Service :  Sign Up - SUCCESS")
	return &pb.SignUpResponse{
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
