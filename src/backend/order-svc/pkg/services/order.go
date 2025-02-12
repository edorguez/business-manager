package services

import (
	"context"
	"fmt"
	"net/http"

	customer "github.com/EdoRguez/business-manager/customer-svc/pkg/pb/customer"
	"github.com/EdoRguez/business-manager/order-svc/pkg/client"
	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	order "github.com/EdoRguez/business-manager/order-svc/pkg/pb/order"
	"github.com/EdoRguez/business-manager/whatsapp-svc/pkg/pb/whatsapp"
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

	_, err := client.SendOrderBusinessMessage(&whatsapp.SendOrderBusinessMessageRequest{ToPhone: req.Customer.Phone, CustomerName: req.Customer.FirstName, ContactNumber: req.Customer.Phone, Products: orderProducts}, ctx)
	if err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	_, err = client.SendOrderCustomerMessage(&whatsapp.SendOrderCustomerMessageRequest{ToPhone: req.Customer.Phone, CustomerName: req.Customer.FirstName, ContactNumber: req.Customer.Phone, Products: orderProducts}, ctx)
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
