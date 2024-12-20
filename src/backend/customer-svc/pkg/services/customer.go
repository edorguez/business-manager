package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/customer-svc/pkg/db/sqlc"
	"github.com/EdoRguez/business-manager/customer-svc/pkg/pb/customer"
	repo "github.com/EdoRguez/business-manager/customer-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/customer-svc/pkg/util/type_converter"
	"google.golang.org/protobuf/types/known/timestamppb"
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

	existCustomer, errExist := s.Repo.GetCustomerByIdentification(ctx, db.GetCustomerByIdentificationParams{
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	})

	if errExist != nil && errExist != sql.ErrNoRows {
		fmt.Println("Customer Service :  CreateCustomer - ERROR")
		fmt.Println(errExist.Error())
		return &customer.CreateCustomerResponse{
			Status: http.StatusConflict,
			Error:  errExist.Error(),
		}, nil
	}

	if existCustomer.ID != 0 {
		fmt.Println("Customer Service :  CreateCustomer - ERROR")
		fmt.Println("Customer already exists")
		return &customer.CreateCustomerResponse{
			Status: http.StatusConflict,
			Error:  "Customer already exists",
		}, nil
	}

	createCustomerParams := db.CreateCustomerParams{
		CompanyID:            req.CompanyId,
		FirstName:            req.FirstName,
		LastName:             type_converter.NewSqlNullString(req.LastName),
		Email:                type_converter.NewSqlNullString(req.Email),
		Phone:                type_converter.NewSqlNullString(req.Phone),
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	}

	c, err := s.Repo.CreateCustomer(ctx, createCustomerParams)
	if err != nil {
		fmt.Println("Customer Service :  CreateCustomer - ERROR")
		fmt.Println(err.Error())
		return &customer.CreateCustomerResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Customer Service :  CreateCustomer - SUCCESS")
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
		fmt.Println("Customer Service :  GetCustomer - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &customer.GetCustomerResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Customer Service :  GetCustomer - SUCCESS")
	return &customer.GetCustomerResponse{
		Id:                   c.ID,
		CompanyId:            c.CompanyID,
		FirstName:            c.FirstName,
		LastName:             type_converter.NewString(c.LastName),
		Email:                type_converter.NewString(c.Email),
		Phone:                type_converter.NewString(c.Phone),
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
		Limit:                req.Limit,
		Offset:               req.Offset,
		CompanyID:            req.CompanyId,
		FirstName:            *req.FirstName,
		LastName:             *req.LastName,
		IdentificationNumber: *req.IdentificationNumber,
	}

	c, err := s.Repo.GetCustomers(ctx, params)
	if err != nil {
		fmt.Println("Customer Service :  GetCustomers - ERROR")
		fmt.Println(err.Error())
		return &customer.GetCustomersResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	customers := make([]*customer.GetCustomerResponse, 0, len(c))
	for _, v := range c {
		customers = append(customers, &customer.GetCustomerResponse{
			Id:                   v.ID,
			CompanyId:            v.CompanyID,
			FirstName:            v.FirstName,
			LastName:             type_converter.NewString(v.LastName),
			Email:                type_converter.NewString(v.Email),
			Phone:                type_converter.NewString(v.Phone),
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
			Status:               http.StatusOK,
		})
	}

	fmt.Println("Customer Service :  GetCustomers - SUCCESS")
	return &customer.GetCustomersResponse{
		Customers: customers,
		Status:    http.StatusOK,
	}, nil
}

func (s *CustomerService) GetCustomersByMonths(ctx context.Context, req *customer.GetCustomersByMonthsRequest) (*customer.GetCustomersByMonthsResponse, error) {
	fmt.Println("Customer Service :  GetCustomersByMonths")
	fmt.Println("Customer Service :  GetCustomersByMonths - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetCustomersByMonthsParams{
		CompanyID: req.CompanyId,
		Months:    req.Months,
	}

	c, err := s.Repo.GetCustomersByMonths(ctx, params)
	if err != nil {
		fmt.Println("Customer Service :  GetCustomersByMonths - ERROR")
		fmt.Println(err.Error())
		return &customer.GetCustomersByMonthsResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	customers := make([]*timestamppb.Timestamp, 0, len(c))
	for _, v := range c {
		customers = append(customers, timestamppb.New(v))
	}

	fmt.Println("Customer Service :  GetCustomersByMonths - SUCCESS")
	return &customer.GetCustomersByMonthsResponse{
		CreatedAt: customers,
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
		LastName:             type_converter.NewSqlNullString(req.LastName),
		Email:                type_converter.NewSqlNullString(req.Email),
		Phone:                type_converter.NewSqlNullString(req.Phone),
		IdentificationNumber: req.IdentificationNumber,
		IdentificationType:   req.IdentificationType,
	}

	_, err := s.Repo.UpdateCustomer(ctx, params)
	if err != nil {
		fmt.Println("Customer Service :  UpdateCustomer - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}
		return &customer.UpdateCustomerResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Customer Service :  UpdateCustomer - SUCCESS")
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
		fmt.Println("Customer Service :  DeleteCustomer - ERROR")
		fmt.Println(err.Error())
		return &customer.DeleteCustomerResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Customer Service :  DeleteCustomer - SUCCESS")
	return &customer.DeleteCustomerResponse{
		Status: http.StatusNoContent,
	}, nil
}
