package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/customer-svc/pkg/db/sqlc"
)

type CustomerRepo struct {
	SQLStorage *db.SQLStorage
}

func NewCustomerRepo(sql *db.SQLStorage) *CustomerRepo {
	return &CustomerRepo{
		SQLStorage: sql,
	}
}

func (clientRepo *CustomerRepo) CreateCustomer(ctx context.Context, arg db.CreateCustomerParams) (db.CustomerCustomer, error) {
	var result db.CustomerCustomer

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateCustomer(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *CustomerRepo) GetCustomer(ctx context.Context, id int64) (db.CustomerCustomer, error) {
	var result db.CustomerCustomer

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCustomer(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *CustomerRepo) GetCustomers(ctx context.Context, arg db.GetCustomersParams) ([]db.CustomerCustomer, error) {
	var result []db.CustomerCustomer

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCustomers(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *CustomerRepo) UpdateCustomer(ctx context.Context, arg db.UpdateCustomerParams) (db.CustomerCustomer, error) {
	var result db.CustomerCustomer

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdateCustomer(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *CustomerRepo) DeleteCustomer(ctx context.Context, id int64) error {
	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		errDelete := q.DeleteCustomer(ctx, id)
		if errDelete != nil {
			return errDelete
		}

		return errDelete
	})

	return err
}

func (clientRepo *CustomerRepo) GetCustomersByMonths(ctx context.Context, companyID int64) ([]db.GetCustomersByMonthsRow, error) {
	var result []db.GetCustomersByMonthsRow

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCustomersByMonths(ctx, companyID)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
