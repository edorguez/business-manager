package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/company-svc/pkg/db/sqlc"
)

type PaymentRepo struct {
	SQLStorage *db.SQLStorage
}

func NewPaymentRepo(sql *db.SQLStorage) *PaymentRepo {
	return &PaymentRepo{
		SQLStorage: sql,
	}
}

func (paymentRepo *PaymentRepo) CreatePayment(ctx context.Context, arg db.CreatePaymentParams) (db.CompanyPayment, error) {
	var result db.CompanyPayment

	err := paymentRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreatePayment(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (paymentRepo *PaymentRepo) GetPayment(ctx context.Context, id int64) (db.GetPaymentRow, error) {
	var result db.GetPaymentRow

	err := paymentRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetPayment(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (paymentRepo *PaymentRepo) GetPayments(ctx context.Context, arg db.GetPaymentsParams) ([]db.GetPaymentsRow, error) {
	var result []db.GetPaymentsRow

	err := paymentRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetPayments(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (paymentRepo *PaymentRepo) UpdatePayment(ctx context.Context, arg db.UpdatePaymentParams) (db.CompanyPayment, error) {
	var result db.CompanyPayment

	err := paymentRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdatePayment(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (paymentRepo *PaymentRepo) DeletePayment(ctx context.Context, id int64) error {
	err := paymentRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		errDelete := q.DeletePayment(ctx, id)
		if errDelete != nil {
			return errDelete
		}

		return errDelete
	})

	return err
}
