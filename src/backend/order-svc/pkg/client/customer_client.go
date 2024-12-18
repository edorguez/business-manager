package client

import (
	"context"
	"fmt"

	pb "github.com/EdoRguez/business-manager/customer-svc/pkg/pb/customer"
	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	"google.golang.org/grpc"
)

var customerServiceClient pb.CustomerServiceClient

func InitCustomerServiceClient(c *config.Config) error {
	fmt.Println("Customer CLIENT :  InitCustomerServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Customer_Svc_Url, grpc.WithInsecure())

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
