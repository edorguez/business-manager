package services

import (
	"context"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/customer-svc/pkg/db/sqlc"
	customer "github.com/EdoRguez/business-manager/customer-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/customer-svc/pkg/repository"
	util "github.com/EdoRguez/business-manager/customer-svc/pkg/util"
)

type CustomerService struct {
	Repo *repo.CustomerRepo
	customer.UnimplementedCustomerServiceServer
}

func (s *CustomerService) CreateCustomer(ctx context.Context, req *customer.CreateCustomerRequest) (*customer.CreateCustomerResponse, error) {
	fmt.Println("Customer Service :  CreateCustomer")
	fmt.Println("Customer Service :  CreateCustomer - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createCustomerParams := db.CreateCustomerParams{
		CompanyID:            req.CompanyId,
		FirstName:            req.FirstName,
		LastName:             util.NewSqlNullString(req.LastName),
		Email:                util.NewSqlNullString(req.Email),
		Phone:                util.NewSqlNullString(req.Phone),
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	}

	c, err := s.Repo.CreateCustomer(ctx, createCustomerParams)
	if err != nil {
		fmt.Println("API Gateway :  CreateCustomer - ERROR")
		fmt.Println(err.Error())
		return &customer.CreateCustomerResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  CreateCustomer - SUCCESS")
	return &customer.CreateCustomerResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *CustomerService) GetCustomer(ctx context.Context, req *customer.GetCustomerRequest) (*customer.GetCustomerResponse, error) {
	fmt.Println("Customer Service :  GetCustomer")
	fmt.Println("Customer Service :  GetCustomer - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetCustomer(ctx, req.Id)
	if err != nil {
		fmt.Println("API Gateway :  GetCustomer - ERROR")
		fmt.Println(err.Error())
		return &customer.GetCustomerResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  GetCustomer - SUCCESS")
	return &customer.GetCustomerResponse{
		Id:                   c.ID,
		CompanyId:            c.CompanyID,
		FirstName:            c.FirstName,
		LastName:             c.LastName.String,
		Email:                c.Email.String,
		Phone:                c.Phone.String,
		IdentificationNumber: c.IdentificationNumber,
		IdentificationType:   c.IdentificationType,
		Status:               http.StatusOK,
	}, nil
}

func (s *CustomerService) GetCustomers(ctx context.Context, req *customer.GetCustomersRequest) (*customer.GetCustomersResponse, error) {
	fmt.Println("Customer Service :  GetCustomers")
	fmt.Println("Customer Service :  GetCustomers - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetCustomersParams{
		CompanyID: req.CompanyId,
		Limit:     req.Limit,
		Offset:    req.Offset,
	}

	c, err := s.Repo.GetCustomers(ctx, params)
	if err != nil {
		fmt.Println("API Gateway :  GetCustomers - ERROR")
		fmt.Println(err.Error())
		return &customer.GetCustomersResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	var customers []*customer.GetCustomerResponse
	for _, v := range c {
		customers = append(customers, &customer.GetCustomerResponse{
			Id:                   v.ID,
			CompanyId:            v.CompanyID,
			FirstName:            v.FirstName,
			LastName:             v.LastName.String,
			Email:                v.Email.String,
			Phone:                v.Phone.String,
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
			Status:               http.StatusOK,
		})
	}

	fmt.Println("API Gateway :  GetCustomers - SUCCESS")
	return &customer.GetCustomersResponse{
		Customers: customers,
		Status:    http.StatusOK,
	}, nil
}

func (s *CustomerService) UpdateCustomer(ctx context.Context, req *customer.UpdateCustomerRequest) (*customer.UpdateCustomerResponse, error) {
	fmt.Println("Customer Service :  UpdateCustomer")
	fmt.Println("Customer Service :  UpdateCustomer - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdateCustomerParams{
		ID:                   req.Id,
		FirstName:            req.FirstName,
		LastName:             util.NewSqlNullString(req.LastName),
		Email:                util.NewSqlNullString(req.Email),
		Phone:                util.NewSqlNullString(req.Phone),
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	}

	_, err := s.Repo.UpdateCustomer(ctx, params)
	if err != nil {
		fmt.Println("API Gateway :  UpdateCustomer - ERROR")
		fmt.Println(err.Error())
		return &customer.UpdateCustomerResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  UpdateCustomer - SUCCESS")
	return &customer.UpdateCustomerResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *CustomerService) DeleteCustomer(ctx context.Context, req *customer.DeleteCustomerRequest) (*customer.DeleteCustomerResponse, error) {
	fmt.Println("Customer Service :  DeleteCustomer")
	fmt.Println("Customer Service :  DeleteCustomer - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeleteCustomer(ctx, req.Id)
	if err != nil {
		fmt.Println("API Gateway :  DeleteCustomer - ERROR")
		fmt.Println(err.Error())
		return &customer.DeleteCustomerResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("API Gateway :  DeleteCustomer - SUCCESS")
	return &customer.DeleteCustomerResponse{
		Status: http.StatusNoContent,
	}, nil
}
