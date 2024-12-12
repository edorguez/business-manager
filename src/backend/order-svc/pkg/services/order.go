package services

import (
	"context"

	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	product "github.com/EdoRguez/business-manager/product-svc/pkg/pb"
)

type ProductService struct {
	Config *config.Config
	order.UnimplementedOrderServiceServer
}

func (s *ProductService) CreateOrder(ctx context.Context, req *product.CreateProductRequest) (*product.CreateProductResponse, error) {

}
