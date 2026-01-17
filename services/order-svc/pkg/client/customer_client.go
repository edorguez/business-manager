package client

import (
	"context"
	"fmt"
	"os"

	"github.com/edorguez/business-manager/services/order-svc/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/customer"
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
	if appEnv == "development" {
		fmt.Println("Running in development mode")
		customerSvcUrl = "localhost:" + c.CustomerSvcPort
	} else {
		fmt.Println("Running in docker mode")
		customerSvcUrl = c.CustomerSvcUrl + ":" + c.CustomerSvcPort
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

func CreateCustomer(params *pb.CreateCustomerRequest, c context.Context) (*pb.CreateCustomerResponse, error) {
	fmt.Println("Customer CLIENT :  CreateCustomer")

	fmt.Println("Customer CLIENT :  CreateCustomer - Body")
	fmt.Println(params)
	fmt.Println("-----------------")

	res, err := customerServiceClient.CreateCustomer(c, params)

	if err != nil {
		fmt.Println("Customer CLIENT :  CreateCustomer - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Customer CLIENT :  CreateCustomer - SUCCESS")
	return res, nil
}

func GetCustomer(params *pb.GetCustomerRequest, c context.Context) (*pb.GetCustomerResponse, error) {
	fmt.Println("Customer CLIENT :  GetCustomer")

	fmt.Println("Customer CLIENT :  GetCustomer - Body")
	fmt.Println(params)
	fmt.Println("-----------------")

	res, err := customerServiceClient.GetCustomer(c, params)

	if err != nil {
		fmt.Println("Customer CLIENT :  GetCustomer - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Customer CLIENT :  GetCustomer - SUCCESS")
	return res, nil
}
