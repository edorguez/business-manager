package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/order/contracts"
	pb "github.com/edorguez/business-manager/shared/pb/order"
	"github.com/edorguez/business-manager/shared/types"
	"google.golang.org/grpc"
)

var orderServiceClient pb.OrderServiceClient

func InitOrderServiceClient(c *config.Config) error {
	fmt.Println("Order CLIENT :  InitOrderServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var orderSvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		orderSvcUrl = c.OrderSvcUrl + ":" + c.OrderSvcPort
	} else {
		fmt.Println("Running in development mode")
		orderSvcUrl = c.DevelopmentUrl + ":" + c.OrderSvcPort
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(orderSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	orderServiceClient = pb.NewOrderServiceClient(cc)
	return nil
}

func CreateOrder(body contracts.CreateOrderRequest, c context.Context) (*pb.CreateOrderResponse, *types.Error) {
	fmt.Println("Order CLIENT :  CreateOrder")

	fmt.Println("Order CLIENT :  CreateOrder - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createOrderParams := &pb.CreateOrderRequest{
		CompanyId: body.CompanyId,
		Customer: &pb.CreateOrderCustomerRequest{
			FirstName:            body.Customer.FirstName,
			LastName:             body.Customer.LastName,
			Phone:                body.Customer.Phone,
			IdentificationNumber: body.Customer.IdentificationNumber,
			IdentificationType:   body.Customer.IdentificationType,
		},
		Products: make([]*pb.CreateOrderProductRequest, 0, len(body.Products)),
	}

	for _, product := range body.Products {
		createOrderParams.Products = append(createOrderParams.Products, &pb.CreateOrderProductRequest{
			ProductId: product.ProductId,
			Name:      product.Name,
			Quantity:  product.Quantity,
			Price:     product.Price,
		})
	}

	res, err := orderServiceClient.CreateOrder(c, createOrderParams)

	if err != nil {
		fmt.Println("Order CLIENT :  CreateOrder - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Order CLIENT :  CreateOrder - SUCCESS")
	return res, nil
}
