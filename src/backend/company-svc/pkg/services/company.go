package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/company-svc/pkg/db/sqlc"
	company "github.com/EdoRguez/business-manager/company-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/company-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/company-svc/pkg/util/type_converter"
)

type CompanyService struct {
	Repo *repo.CompanyRepo
	company.UnimplementedCompanyServiceServer
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *company.CreateCompanyRequest) (*company.CreateCompanyResponse, error) {
	fmt.Println("Company Service :  CreateCompany")
	fmt.Println("Company Service :  CreateCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createCompanyParams := db.CreateCompanyParams{
		Name:     req.Name,
		ImageUrl: type_converter.NewSqlNullString(req.ImageUrl),
	}

	c, err := s.Repo.CreateCompany(ctx, createCompanyParams)
	if err != nil {
		fmt.Println("Company Service :  CreateCompany - ERROR")
		fmt.Println(err.Error())
		return &company.CreateCompanyResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Company Service :  CreateCompany - SUCCESS")
	return &company.CreateCompanyResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *CompanyService) GetCompany(ctx context.Context, req *company.GetCompanyRequest) (*company.GetCompanyResponse, error) {
	fmt.Println("Company Service :  GetCompany")
	fmt.Println("Company Service :  GetCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetCompany(ctx, req.Id)
	if err != nil {
		fmt.Println("API Gateway :  GetCompany - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.GetCompanyResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  GetCompany - SUCCESS")
	return &company.GetCompanyResponse{
		Id:       c.ID,
		Name:     c.Name,
		ImageUrl: type_converter.NewString(c.ImageUrl),
		Status:   http.StatusOK,
	}, nil
}

func (s *CompanyService) GetCompanies(ctx context.Context, req *company.GetCompaniesRequest) (*company.GetCompaniesResponse, error) {
	fmt.Println("Company Service :  GetCompanies")
	fmt.Println("Company Service :  GetCompanies - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetCompaniesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	c, err := s.Repo.GetCompanies(ctx, params)
	if err != nil {
		fmt.Println("Company Service :  GetCompanies - ERROR")
		fmt.Println(err.Error())
		return &company.GetCompaniesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	var companies []*company.GetCompanyResponse
	for _, v := range c {
		companies = append(companies, &company.GetCompanyResponse{
			Id:       v.ID,
			Name:     v.Name,
			ImageUrl: type_converter.NewString(v.ImageUrl),
			Status:   http.StatusOK,
		})
	}

	fmt.Println("Company Service :  GetCompanies - SUCCESS")
	return &company.GetCompaniesResponse{
		Companies: companies,
		Status:    http.StatusOK,
	}, nil
}

func (s *CompanyService) UpdateCompany(ctx context.Context, req *company.UpdateCompanyRequest) (*company.UpdateCompanyResponse, error) {
	fmt.Println("Company Service :  UpdateCompany")
	fmt.Println("Company Service :  UpdateCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdateCompanyParams{
		ID:       req.Id,
		Name:     req.Name,
		ImageUrl: type_converter.NewSqlNullString(req.ImageUrl),
	}

	_, err := s.Repo.UpdateCompany(ctx, params)
	if err != nil {
		fmt.Println("Company Service :  UpdateCompany - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.UpdateCompanyResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  UpdateCompany - SUCCESS")
	return &company.UpdateCompanyResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *CompanyService) DeleteCompany(ctx context.Context, req *company.DeleteCompanyRequest) (*company.DeleteCompanyResponse, error) {
	fmt.Println("Company Service :  DeleteCompany")
	fmt.Println("Company Service :  DeleteCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeleteCompany(ctx, req.Id)
	if err != nil {
		fmt.Println("Company Service :  DeleteCompany - ERROR")
		fmt.Println(err.Error())
		return &company.DeleteCompanyResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Company Service :  DeleteCompany - SUCCESS")
	return &company.DeleteCompanyResponse{
		Status: http.StatusNoContent,
	}, nil
}
