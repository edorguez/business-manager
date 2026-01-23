package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/order-svc/pkg/client"
	"github.com/edorguez/business-manager/services/order-svc/pkg/config"
	"github.com/edorguez/business-manager/services/order-svc/pkg/repository"
	customer "github.com/edorguez/business-manager/shared/pb/customer"
	order "github.com/edorguez/business-manager/shared/pb/order"
	product "github.com/edorguez/business-manager/shared/pb/product"
	"github.com/edorguez/business-manager/shared/pb/whatsapp"
	"github.com/edorguez/business-manager/shared/util/type_converter"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type OrderService struct {
	Config *config.Config
	Repo   *repository.OrderRepo
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

	productClientErr := client.InitProductServiceClient(s.Config)
	if productClientErr != nil {
		fmt.Println("Order Service :  CreateOrder - WARNING: failed to init product service client")
		fmt.Println(productClientErr.Error())
	}

	whatsappProducts := make([]*whatsapp.OrderProductRequest, 0, len(req.Products))
	for _, v := range req.Products {
		whatsappProducts = append(whatsappProducts, &whatsapp.OrderProductRequest{
			Name:     v.Name,
			Quantity: v.Quantity,
			Price:    v.Price,
		})
	}

	_, err := client.SendOrderBusinessMessage(&whatsapp.SendOrderBusinessMessageRequest{CompanyId: req.CompanyId, CustomerName: req.Customer.FirstName, ContactNumber: req.Customer.Phone, Products: whatsappProducts}, ctx)
	if err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	_, err = client.SendOrderCustomerMessage(&whatsapp.SendOrderCustomerMessageRequest{CompanyId: req.CompanyId, ToPhone: req.Customer.Phone, CustomerName: req.Customer.FirstName, Products: whatsappProducts}, ctx)
	if err != nil {
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Create customer first
	customerResp, err := client.CreateCustomer(&customer.CreateCustomerRequest{
		CompanyId:            req.CompanyId,
		FirstName:            req.Customer.FirstName,
		LastName:             req.Customer.LastName,
		Phone:                &req.Customer.Phone,
		IdentificationNumber: req.Customer.IdentificationNumber,
		IdentificationType:   req.Customer.IdentificationType,
	}, ctx)

	if err != nil {
		fmt.Println("Order Service :  CreateOrder - ERROR creating customer")
		fmt.Println(err.Error())
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	if customerResp.Status != http.StatusCreated && customerResp.Status != http.StatusOK {
		fmt.Println("Order Service :  CreateOrder - Customer service returned error")
		return &order.CreateOrderResponse{
			Status: customerResp.Status,
			Error:  customerResp.Error,
		}, nil
	}

	// Create order with products in database
	// First, fetch product images from product service if available
	productImageMap := make(map[string]*string) // product ID -> first image URL (nil if none)
	if productClientErr == nil {
		// Collect product IDs
		productIds := make([]string, 0, len(req.Products))
		for _, v := range req.Products {
			productIds = append(productIds, v.ProductId)
		}

		// Call product service to get product details
		productResp, err := client.GetProductsByIds(&product.GetProductsByIdsRequest{
			CompanyId: req.CompanyId,
			Ids:       productIds,
		}, ctx)

		if err != nil {
			fmt.Println("Order Service :  CreateOrder - WARNING: failed to fetch product images")
			fmt.Println(err.Error())
			// Continue with empty image map
		} else if productResp.Status != http.StatusOK {
			fmt.Println("Order Service :  CreateOrder - WARNING: product service returned error status:", productResp.Status, productResp.Error)
			// Continue with empty image map
		} else {
			// Build map of product ID to first image URL
			for _, p := range productResp.Products {
				var imageUrl *string
				if len(p.Images) > 0 {
					imageUrl = &p.Images[0]
				}
				productImageMap[p.Id] = imageUrl
			}
		}
	}

	dbOrderProducts := make([]repository.CreateOrderProductParams, 0, len(req.Products))
	for _, v := range req.Products {
		var imageUrl *string
		if img, exists := productImageMap[v.ProductId]; exists {
			imageUrl = img
		}
		dbOrderProducts = append(dbOrderProducts, repository.CreateOrderProductParams{
			ProductID: v.ProductId,
			Name:      v.Name,
			Quantity:  v.Quantity,
			Price:     v.Price,
			ImageUrl:  imageUrl,
		})
	}

	params := repository.CreateOrderWithProductsParams{
		CompanyID:  req.CompanyId,
		CustomerID: customerResp.Id,
		Products:   dbOrderProducts,
	}

	_, err = s.Repo.CreateOrderWithProducts(ctx, params)
	if err != nil {
		fmt.Println("Order Service :  CreateOrder - ERROR creating order in database")
		fmt.Println(err.Error())
		return &order.CreateOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Order Service :  CreateOrder - SUCCESS")
	return &order.CreateOrderResponse{
		Status: http.StatusCreated,
	}, nil
}

func (s *OrderService) GetOrder(ctx context.Context, req *order.GetOrderRequest) (*order.GetOrderResponse, error) {
	fmt.Println("Order Service :  GetOrder")
	fmt.Println("Order Service :  GetOrder - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	// Parse order ID from string to int64
	var orderID int64
	_, err := fmt.Sscanf(req.Id, "%d", &orderID)
	if err != nil {
		return &order.GetOrderResponse{
			Status: http.StatusBadRequest,
			Error:  "Invalid order ID format",
		}, nil
	}

	// Fetch order from repository
	dbOrder, err := s.Repo.GetOrder(ctx, orderID)
	if err != nil {
		return &order.GetOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Initialize customer service client
	if err := client.InitCustomerServiceClient(s.Config); err != nil {
		return &order.GetOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Fetch customer details
	customerResp, err := client.GetCustomer(&customer.GetCustomerRequest{
		Id: dbOrder.CustomerID,
	}, ctx)
	if err != nil {
		return &order.GetOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	if customerResp.Status != http.StatusOK {
		return &order.GetOrderResponse{
			Status: customerResp.Status,
			Error:  customerResp.Error,
		}, nil
	}

	// Fetch order products
	products, err := s.Repo.GetOrderProductsByOrderId(ctx, orderID)
	if err != nil {
		return &order.GetOrderResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Convert database order to protobuf order
	orderProto := &order.Order{
		Id:          dbOrder.ID,
		CompanyId:   dbOrder.CompanyID,
		CustomerId:  dbOrder.CustomerID,
		CreatedAt:   timestamppb.New(dbOrder.CreatedAt),
		OrderNumber: int64(dbOrder.OrderNumber),
	}

	// Convert customer response to protobuf customer
	customerProto := &order.Customer{
		Id:                   customerResp.Id,
		FirstName:            customerResp.FirstName,
		LastName:             customerResp.LastName,
		Email:                customerResp.Email,
		Phone:                customerResp.Phone,
		IdentificationNumber: customerResp.IdentificationNumber,
		IdentificationType:   customerResp.IdentificationType,
	}

	// Convert products to protobuf products
	var productsProto []*order.OrderProduct
	for _, p := range products {
		productsProto = append(productsProto, &order.OrderProduct{
			Id:        p.ID,
			OrderId:   p.OrderID,
			ProductId: p.ProductID,
			Name:      p.Name,
			Quantity:  uint32(p.Quantity),
			Price:     uint64(p.Price),
			ImageUrl:  type_converter.NewString(p.ImageUrl),
		})
	}

	fmt.Println("Order Service :  GetOrder - SUCCESS")
	return &order.GetOrderResponse{
		Order:    orderProto,
		Customer: customerProto,
		Products: productsProto,
		Status:   http.StatusOK,
	}, nil
}

func (s *OrderService) GetOrders(ctx context.Context, req *order.GetOrdersRequest) (*order.GetOrdersResponse, error) {
	fmt.Println("Order Service :  GetOrders")
	fmt.Println("Order Service :  GetOrders - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	// Initialize customer service client
	if err := client.InitCustomerServiceClient(s.Config); err != nil {
		return &order.GetOrdersResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Fetch orders from repository
	orders, err := s.Repo.GetOrders(ctx, req.CompanyId, req.Limit, req.Offset)
	if err != nil {
		return &order.GetOrdersResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Get total count for pagination
	total, err := s.Repo.GetOrdersCount(ctx, req.CompanyId)
	if err != nil {
		return &order.GetOrdersResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	// Prepare response
	var ordersWithDetails []*order.OrderWithDetails

	// Process each order
	for _, dbOrder := range orders {
		// Fetch customer details
		customerResp, err := client.GetCustomer(&customer.GetCustomerRequest{
			Id: dbOrder.CustomerID,
		}, ctx)
		if err != nil {
			// If customer fetch fails, we might want to skip this order or return error
			// For now, skip this order and continue
			continue
		}

		if customerResp.Status != http.StatusOK {
			// Skip this order if customer not found
			continue
		}

		// Fetch order products
		products, err := s.Repo.GetOrderProductsByOrderId(ctx, dbOrder.ID)
		if err != nil {
			// Skip this order if products fetch fails
			continue
		}

		// Convert database order to protobuf order
		orderProto := &order.Order{
			Id:          dbOrder.ID,
			CompanyId:   dbOrder.CompanyID,
			CustomerId:  dbOrder.CustomerID,
			CreatedAt:   timestamppb.New(dbOrder.CreatedAt),
			OrderNumber: int64(dbOrder.OrderNumber),
		}

		// Convert customer response to protobuf customer
		customerProto := &order.Customer{
			Id:                   customerResp.Id,
			FirstName:            customerResp.FirstName,
			LastName:             customerResp.LastName,
			Email:                customerResp.Email,
			Phone:                customerResp.Phone,
			IdentificationNumber: customerResp.IdentificationNumber,
			IdentificationType:   customerResp.IdentificationType,
		}

		// Convert products to protobuf products
		var productsProto []*order.OrderProduct
		for _, p := range products {
			productsProto = append(productsProto, &order.OrderProduct{
				Id:        p.ID,
				OrderId:   p.OrderID,
				ProductId: p.ProductID,
				Name:      p.Name,
				Quantity:  uint32(p.Quantity),
				Price:     uint64(p.Price),
				ImageUrl:  type_converter.NewString(p.ImageUrl),
			})
		}

		// Create order with details
		ordersWithDetails = append(ordersWithDetails, &order.OrderWithDetails{
			Order:    orderProto,
			Customer: customerProto,
			Products: productsProto,
		})
	}

	fmt.Println("Order Service :  GetOrders - SUCCESS")
	return &order.GetOrdersResponse{
		Orders: ordersWithDetails,
		Total:  total,
		Status: http.StatusOK,
	}, nil
}
