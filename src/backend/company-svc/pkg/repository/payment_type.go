package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/company-svc/pkg/db/sqlc"
)

type PaymentTypeRepo struct {
	SQLStorage *db.SQLStorage
}

func NewPaymentTypeRepo(sql *db.SQLStorage) *PaymentTypeRepo {
	return &PaymentTypeRepo{
		SQLStorage: sql,
	}
}

func (paymentTypeRepo *PaymentTypeRepo) GetPaymentType(ctx context.Context, id int64) (db.CompanyPaymentType, error) {
	var result db.CompanyPaymentType

	err := paymentTypeRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetPaymentType(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (paymentTypeRepo *PaymentTypeRepo) GetPaymentTypes(ctx context.Context, arg db.GetPaymentTypesParams) ([]db.CompanyPaymentType, error) {
	var result []db.CompanyPaymentType

	err := paymentTypeRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetPaymentTypes(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
