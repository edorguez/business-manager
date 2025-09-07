package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/product/contracts"
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
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		productSvcUrl = c.Product_Svc_Url + ":" + c.Product_Svc_Port
	} else {
		fmt.Println("Running in development mode")
		productSvcUrl = c.Development_Url + ":" + c.Product_Svc_Port
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

func CreateProduct(body contracts.CreateProductRequest, images [][]byte, c context.Context) (*pb.CreateProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  CreateProduct")

	fmt.Println("Product CLIENT :  CreateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createProductParams := &pb.CreateProductRequest{
		CompanyId:     body.CompanyId,
		Name:          body.Name,
		Description:   body.Description,
		Sku:           body.Sku,
		Quantity:      body.Quantity,
		Price:         body.Price,
		Images:        images,
		ProductStatus: body.ProductStatus,
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

	fmt.Println("Product CLIENT :  CreateProduct - SUCCESS")
	return res, nil
}

func GetProduct(id string, c context.Context) (*contracts.GetProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  GetProduct")

	params := &pb.GetProductRequest{
		Id: id,
	}

	res, err := productServiceClient.GetProduct(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  GetProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	fmt.Println("Product CLIENT :  GetProduct - SUCCESS")
	return &contracts.GetProductResponse{
		Id:            res.Id,
		CompanyId:     res.CompanyId,
		Name:          res.Name,
		Description:   res.Description,
		Sku:           res.Sku,
		Quantity:      res.Quantity,
		Price:         res.Price,
		Images:        res.Images,
		ProductStatus: res.ProductStatus,
	}, nil
}

func GetProducts(params *pb.GetProductsRequest, c context.Context) ([]*contracts.GetProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  GetProduct")

	if params.CompanyId <= 0 {
		error := &contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		}

		return nil, error
	}

	res, err := productServiceClient.GetProducts(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  GetProducts - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	pr := make([]*contracts.GetProductResponse, 0, len(res.Products))
	for _, v := range res.Products {
		pr = append(pr, &contracts.GetProductResponse{
			Id:            v.Id,
			CompanyId:     v.CompanyId,
			Name:          v.Name,
			Description:   v.Description,
			Sku:           v.Sku,
			Quantity:      v.Quantity,
			Price:         v.Price,
			Images:        v.Images,
			ProductStatus: v.ProductStatus,
		})
	}

	fmt.Println("Product CLIENT :  GetProducts - SUCCESS")
	return pr, nil
}

func GetLatestProducts(params *pb.GetLatestProductsRequest, c context.Context) ([]*contracts.GetProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  GetLatestProduct")

	if params.CompanyId <= 0 {
		error := &contracts.Error{
			Status: http.StatusBadRequest,
			Error:  "Company ID is required in order to get results",
		}

		return nil, error
	}

	res, err := productServiceClient.GetLatestProducts(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  GetLatestProducts - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	pr := make([]*contracts.GetProductResponse, 0, len(res.Products))
	for _, v := range res.Products {
		pr = append(pr, &contracts.GetProductResponse{
			Id:            v.Id,
			CompanyId:     v.CompanyId,
			Name:          v.Name,
			Description:   v.Description,
			Sku:           v.Sku,
			Quantity:      v.Quantity,
			Price:         v.Price,
			Images:        v.Images,
			ProductStatus: v.ProductStatus,
		})
	}

	fmt.Println("Product CLIENT :  GetProducts - SUCCESS")
	return pr, nil
}

func UpdateProduct(id string, body contracts.UpdateProductRequest, images [][]byte, c context.Context) (*pb.UpdateProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  UpdateProduct")

	fmt.Println("Product CLIENT :  UpdateProduct - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateProductParams := &pb.UpdateProductRequest{
		Id:          id,
		CompanyId:   body.CompanyId,
		Name:        body.Name,
		Description: body.Description,
		Sku:         body.Sku,
		Quantity:    body.Quantity,
		Price:       body.Price,
		Images:      images,
	}

	res, err := productServiceClient.UpdateProduct(c, updateProductParams)

	if err != nil {
		fmt.Println("Product CLIENT :  UpdateProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Product CLIENT :  UpdateProduct - SUCCESS")
	return res, nil
}

func UpdateProductStatus(id string, body contracts.UpdateProductStatusRequest, c context.Context) (*pb.UpdateProductStatusResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  UpdateProductStatus")

	fmt.Println("Product CLIENT :  UpdateProductStatus - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	var status uint32 = 0
	if body.ProductStatus != nil {
		status = uint32(*body.ProductStatus)
	}

	updateProductParams := &pb.UpdateProductStatusRequest{
		Id:            id,
		ProductStatus: status,
	}

	res, err := productServiceClient.UpdateProductStatus(c, updateProductParams)

	if err != nil {
		fmt.Println("Product CLIENT :  UpdateProductStatus - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Product CLIENT :  UpdateProductStatus - SUCCESS")
	return res, nil
}

func DeleteProduct(id string, c context.Context) (*pb.DeleteProductResponse, *contracts.Error) {
	fmt.Println("Product CLIENT :  DeleteProduct")

	params := &pb.DeleteProductRequest{
		Id: id,
	}

	res, err := productServiceClient.DeleteProduct(c, params)

	if err != nil {
		fmt.Println("Product CLIENT :  DeleteProduct - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Product CLIENT :  DeleteProduct - SUCCESS")
	return res, nil
}
