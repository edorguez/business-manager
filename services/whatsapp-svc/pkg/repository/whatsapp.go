package repository

import (
	"context"

	db "github.com/edorguez/business-manager/services/whatsapp-svc/pkg/db/sqlc"
)

type WhatsappRepo struct {
	SQLStorage *db.SQLStorage
}

func NewWhatsappRepo(sql *db.SQLStorage) *WhatsappRepo {
	return &WhatsappRepo{
		SQLStorage: sql,
	}
}

func (wr *WhatsappRepo) CreateBusinessPhone(ctx context.Context, arg db.CreateBusinessPhoneParams) (db.WhatsappBusinessPhone, error) {
	var result db.WhatsappBusinessPhone

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateBusinessPhone(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (wr *WhatsappRepo) GetBusinessPhoneByCompanyId(ctx context.Context, companyId int64) (db.WhatsappBusinessPhone, error) {
	var result db.WhatsappBusinessPhone

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetBusinessPhoneByCompanyId(ctx, companyId)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (wr *WhatsappRepo) UpdateBusinessPhone(ctx context.Context, arg db.UpdateBusinessPhoneParams) (db.WhatsappBusinessPhone, error) {
	var result db.WhatsappBusinessPhone

	err := wr.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdateBusinessPhone(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
