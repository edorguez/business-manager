package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/product/pb"
	"google.golang.org/grpc"
)

var productServiceClient pb.ProductServiceClient

func InitProductServiceClient(c *config.Config) error {
	fmt.Println("Product CLIENT :  InitProductServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Product_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	productServiceClient = pb.NewProductServiceClient(cc)
	return nil
}

func CreateProduct(body contracts.CreateProductRequest, c context.Context) (*pb.CreateProductResponse, *contracts.Error) {
	fmt.Println("Proudct CLIENT :  CreateProduct")

	fmt.Println("Product CLIENT :  CreateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createProductParams := &pb.CreateProductRequest{
		CompanyId:   body.CompanyId,
		Name:        body.Name,
		Description: body.Description,
		Sku:         body.Sku,
		Price:       body.Price,
	}

	res, err := productServiceClient.CreateProduct(c, createProductParams)

	if err != nil {
		fmt.Println("Product CLIENT :  CreateProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("API Gateway :  CreateProduct - SUCCESS")
	return res, nil
}
