package client

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/customer/contracts"
	pb "github.com/edorguez/business-manager/shared/pb/customer"
	"github.com/edorguez/business-manager/shared/types"
	"google.golang.org/grpc"
)

var customerServiceClient pb.CustomerServiceClient

func InitCustomerServiceClient(c *config.Config) error {
	fmt.Println("Customer CLIENT :  InitCustomerServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var customerSvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		customerSvcUrl = c.CustomerSvcUrl + ":" + c.CustomerSvcPort
	} else {
		fmt.Println("Running in development mode")
		customerSvcUrl = c.DevelopmentUrl + ":" + c.CustomerSvcPort
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(customerSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	customerServiceClient = pb.NewCustomerServiceClient(cc)
	return nil
}

func CreateCustomer(body contracts.CreateCustomerRequest, c context.Context) (*pb.CreateCustomerResponse, *types.Error) {
	fmt.Println("Customer CLIENT :  CreateCustomer")

	fmt.Println("Customer CLIENT :  CreateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createCustomerParams := &pb.CreateCustomerRequest{
		CompanyId:            body.CompanyId,
		FirstName:            body.FirstName,
		LastName:             body.LastName,
		Email:                body.Email,
		Phone:                body.Phone,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
	}

	res, err := customerServiceClient.CreateCustomer(c, createCustomerParams)

	if err != nil {
		fmt.Println("Customer CLIENT :  CreateCustomer - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("API Gateway :  CreateCustomer - SUCCESS")
	return res, nil
}

func GetCustomer(id int64, c context.Context) (*contracts.GetCustomerResponse, *types.Error) {
	fmt.Println("Customer CLIENT :  GetCustomer")

	params := &pb.GetCustomerRequest{
		Id: id,
	}

	res, err := customerServiceClient.GetCustomer(c, params)

	if err != nil {
		fmt.Println("Customer CLIENT :  GetCustomer - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Customer CLIENT :  GetCustomer - SUCCESS")

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetCustomerResponse{
		Id:                   res.Id,
		CompanyId:            res.CompanyId,
		FirstName:            res.FirstName,
		LastName:             res.LastName,
		Email:                res.Email,
		Phone:                res.Phone,
		IdentificationNumber: res.IdentificationNumber,
		IdentificationType:   res.IdentificationType,
	}, nil
}

func GetCustomers(params *pb.GetCustomersRequest, c context.Context) ([]*contracts.GetCustomerResponse, *types.Error) {
	fmt.Println("Customer CLIENT :  GetCustomer")

	if params.CompanyId <= 0 {
		error := &types.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		}

		return nil, error
	}

	res, err := customerServiceClient.GetCustomers(c, params)

	if err != nil {
		fmt.Println("Customer CLIENT :  GetCustomers - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	cr := make([]*contracts.GetCustomerResponse, 0, len(res.Customers))
	for _, v := range res.Customers {
		cr = append(cr, &contracts.GetCustomerResponse{
			Id:                   v.Id,
			CompanyId:            v.CompanyId,
			FirstName:            v.FirstName,
			LastName:             v.LastName,
			Email:                v.Email,
			Phone:                v.Phone,
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
		})
	}

	fmt.Println("Customer CLIENT :  GetCustomers - SUCCESS")
	return cr, nil
}

func GetCustomersByMonths(params *pb.GetCustomersByMonthsRequest, c context.Context) (*contracts.GetCustomerByMonthsResponse, *types.Error) {
	fmt.Println("Customer CLIENT :  GetCustomerByMonths")

	if params.CompanyId <= 0 {
		error := &types.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		}

		return nil, error
	}

	if params.Months <= 0 {
		error := &types.Error{
			Status: http.StatusBadRequest,
			Error:  "Months number is required in order to get results",
		}

		return nil, error
	}

	res, err := customerServiceClient.GetCustomersByMonths(c, params)

	if err != nil {
		fmt.Println("Customer CLIENT :  GetCustomersByMonths - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	cr := make([]time.Time, 0, len(res.CreatedAt))
	for _, v := range res.CreatedAt {
		cr = append(cr, v.AsTime())
	}

	result := &contracts.GetCustomerByMonthsResponse{
		Dates: cr,
	}

	fmt.Println("Customer CLIENT :  GetCustomersByMonths - SUCCESS")
	return result, nil
}

func UpdateCustomer(id int64, body contracts.UpdateCustomerRequest, c context.Context) (*pb.UpdateCustomerResponse, *types.Error) {
	fmt.Println("Customer CLIENT :  UpdateCustomer")

	fmt.Println("Customer CLIENT :  UpdateCustomer - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateCustomerParams := &pb.UpdateCustomerRequest{
		Id:                   id,
		FirstName:            body.FirstName,
		LastName:             body.LastName,
		Email:                body.Email,
		Phone:                body.Phone,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
	}

	res, err := customerServiceClient.UpdateCustomer(c, updateCustomerParams)

	if err != nil {
		fmt.Println("Customer CLIENT :  UpdateCustomer - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Customer CLIENT :  UpdateCustomer - SUCCESS")
	return res, nil
}

func DeleteCustomer(id int64, c context.Context) (*pb.DeleteCustomerResponse, *types.Error) {
	fmt.Println("Customer CLIENT :  DeleteCustomer")

	params := &pb.DeleteCustomerRequest{
		Id: id,
	}

	res, err := customerServiceClient.DeleteCustomer(c, params)

	if err != nil {
		fmt.Println("Customer CLIENT :  DeleteCustomer - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Customer CLIENT :  DeleteCustomer - SUCCESS")
	return res, nil
}
