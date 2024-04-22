package services

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/EdoRguez/business-manager/product-svc/pkg/models"
	product "github.com/EdoRguez/business-manager/product-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/product-svc/pkg/repository"
)

type ProductService struct {
	Repo *repo.ProductRepo
	product.UnimplementedProductServiceServer
}

func (s *ProductService) CreateProduct(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {
	fmt.Println("Product Service :  CreateProduct")
	fmt.Println("Product Service :  CreateProduct - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createProductParams := models.Product{
		CompanyId:   req.CompanyId,
		Name:        req.Name,
		Description: *req.Description,
		Sku:         *req.Sku,
		Price:       req.Price,
		Images:      req.Images,
		Status:      req.Status,
		CreatedAt:   time.Now(),
		ModifiedAt:  time.Now(),
	}

	c, err := s.Repo.CreateProduct(ctx, createProductParams)
	if err != nil {
		fmt.Println("Product Service :  CreateProduct - ERROR")
		fmt.Println(err.Error())
		return &product.CreateProductResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Product Service :  CreateProduct - SUCCESS")
	fmt.Println(c)
	fmt.Println("------------------")

	return &product.CreateProductResponse{
		Status: http.StatusCreated,
		Id:     c.Hex(),
	}, nil
}
