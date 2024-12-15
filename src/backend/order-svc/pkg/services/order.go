package services

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/EdoRguez/business-manager/customer-svc/pkg/pb"
	"github.com/EdoRguez/business-manager/order-svc/pkg/client"
	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	order "github.com/EdoRguez/business-manager/order-svc/pkg/pb"
)

type OrderService struct {
	Config *config.Config
	order.UnimplementedOrderServiceServer
}

func (s *OrderService) CreateOrder(ctx context.Context, req *order.CreateOrderRequest) (*order.CreateOrderResponse, error) {
	fmt.Println("Order Service :  CreateOrder")
	fmt.Println("Order Service :  CreateOrder - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	if err := client.InitCustomerServiceClient(s.Config); err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		_, err := client.CreateCustomer(&pb.CreateCustomerRequest{
			CompanyId:            req.CompanyId,
			FirstName:            req.Customer.FirstName,
			LastName:             req.Customer.LastName,
			Phone:                &req.Customer.Phone,
			IdentificationNumber: req.Customer.IdentificationNumber,
			IdentificationType:   req.Customer.IdentificationType,
		}, ctx)

		if err != nil {
			fmt.Println("Order Service :  CreateOrder - ERROR")
			fmt.Println(err.Error())
		}
	}()

	wg.Wait()

	fmt.Println("Order Service :  CreateOrder - SUCCESS")
	return &order.CreateOrderResponse{
		Status: http.StatusCreated,
	}, nil
}
