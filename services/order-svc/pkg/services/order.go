package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/order-svc/pkg/client"
	"github.com/edorguez/business-manager/services/order-svc/pkg/config"
	customer "github.com/edorguez/business-manager/shared/pb/customer"
	order "github.com/edorguez/business-manager/shared/pb/order"
	"github.com/edorguez/business-manager/shared/pb/whatsapp"
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

	if err := client.InitWhatsappServiceClient(s.Config); err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	orderProducts := make([]*whatsapp.OrderProductRequest, 0, len(req.Products))
	for _, v := range req.Products {
		orderProducts = append(orderProducts, &whatsapp.OrderProductRequest{
			Name:     v.Name,
			Quantity: v.Quantity,
			Price:    v.Price,
		})
	}

	_, err := client.SendOrderBusinessMessage(&whatsapp.SendOrderBusinessMessageRequest{CompanyId: req.CompanyId, CustomerName: req.Customer.FirstName, ContactNumber: req.Customer.Phone, Products: orderProducts}, ctx)
	if err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	_, err = client.SendOrderCustomerMessage(&whatsapp.SendOrderCustomerMessageRequest{CompanyId: req.CompanyId, ToPhone: req.Customer.Phone, CustomerName: req.Customer.FirstName, Products: orderProducts}, ctx)
	if err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	errChan := make(chan error)
	go func() {
		defer close(errChan)
		_, err := client.CreateCustomer(&customer.CreateCustomerRequest{
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
			errChan <- err
			return
		}

		errChan <- nil
	}()

	if errCust := <-errChan; errCust != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  errCust.Error(),
		}, nil
	}

	fmt.Println("Order Service :  CreateOrder - SUCCESS")
	return &order.CreateOrderResponse{
		Status: http.StatusCreated,
	}, nil
}
