package repository

import (
	"context"

	db "github.com/edorguez/business-manager/services/order-svc/pkg/db/sqlc"
	"github.com/edorguez/business-manager/shared/util/type_converter"
)

type OrderRepo struct {
	SQLStorage *db.SQLStorage
}

func NewOrderRepo(sql *db.SQLStorage) *OrderRepo {
	return &OrderRepo{
		SQLStorage: sql,
	}
}

type CreateOrderWithProductsParams struct {
	CompanyID  int64
	CustomerID int64
	Products   []CreateOrderProductParams
}

type CreateOrderProductParams struct {
	ProductID string
	Name      string
	Quantity  uint32
	Price     uint64
	ImageUrl  *string
}

func (repo *OrderRepo) CreateOrderWithProducts(ctx context.Context, arg CreateOrderWithProductsParams) (db.OrderOrder, error) {
	var result db.OrderOrder

	err := repo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		// Acquire advisory lock for this company to prevent concurrent order number assignment
		err = q.LockCompanyOrders(ctx, arg.CompanyID)
		if err != nil {
			return err
		}

		// Get max order number for this company
		maxOrderNumber, err := q.GetMaxOrderNumberForCompany(ctx, arg.CompanyID)
		if err != nil {
			return err
		}
		nextOrderNumber := maxOrderNumber + 1

		// Create order with order number
		orderParams := db.CreateOrderParams{
			CompanyID:   arg.CompanyID,
			CustomerID:  arg.CustomerID,
			OrderNumber: nextOrderNumber,
		}
		result, err = q.CreateOrder(ctx, orderParams)
		if err != nil {
			return err
		}

		// Create order products
		for _, product := range arg.Products {
			productParams := db.CreateOrderProductParams{
				OrderID:   result.ID,
				ProductID: product.ProductID,
				Quantity:  int32(product.Quantity),
				Price:     int64(product.Price),
				Name:      product.Name,
				ImageUrl:  type_converter.NewSqlNullString(product.ImageUrl),
			}
			_, err = q.CreateOrderProduct(ctx, productParams)
			if err != nil {
				return err
			}
		}

		return nil
	})

	return result, err
}

func (repo *OrderRepo) GetOrder(ctx context.Context, id int64) (db.OrderOrder, error) {
	var result db.OrderOrder

	err := repo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetOrder(ctx, id)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func (repo *OrderRepo) GetOrders(ctx context.Context, companyID int64, limit, offset int32) ([]db.OrderOrder, error) {
	var result []db.OrderOrder

	err := repo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		params := db.GetOrdersParams{
			CompanyID: companyID,
			Limit:     limit,
			Offset:    offset,
		}
		result, err = q.GetOrders(ctx, params)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func (repo *OrderRepo) GetOrdersCount(ctx context.Context, companyID int64) (int64, error) {
	var result int64

	err := repo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetOrdersCount(ctx, companyID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}

func (repo *OrderRepo) GetOrderProductsByOrderId(ctx context.Context, orderID int64) ([]db.OrderOrderProduct, error) {
	var result []db.OrderOrderProduct

	err := repo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetOrderProductsByOrderId(ctx, orderID)
		if err != nil {
			return err
		}

		return nil
	})

	return result, err
}
