package client

import (
	"context"
	"fmt"
	"os"

	"github.com/edorguez/business-manager/services/order-svc/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/product"
	"google.golang.org/grpc"
)

var productServiceClient pb.ProductServiceClient

func InitProductServiceClient(c *config.Config) error {
	fmt.Println("Product CLIENT :  InitProductServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var productSvcUrl string
	if appEnv == "development" {
		fmt.Println("Running in development mode")
		productSvcUrl = "localhost:" + c.ProductSvcPort
	} else {
		fmt.Println("Running in docker mode")
		productSvcUrl = c.ProductSvcUrl + ":" + c.ProductSvcPort
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(productSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	productServiceClient = pb.NewProductServiceClient(cc)
	return nil
}

func GetProductsByIds(params *pb.GetProductsByIdsRequest, c context.Context) (*pb.GetProductsByIdsResponse, error) {
	fmt.Println("Product CLIENT :  GetProductsByIds")

	fmt.Println("Product CLIENT :  GetProductsByIds - Body")
	fmt.Println(params)
	fmt.Println("-----------------")

	res, err := productServiceClient.GetProductsByIds(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  GetProductsByIds - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Product CLIENT :  GetProductsByIds - SUCCESS")
	return res, nil
}
