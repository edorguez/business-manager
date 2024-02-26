package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/client-svc/pkg/db/sqlc"
)

type ClientRepo struct {
	SQLStorage *db.SQLStorage
}

func NewClientRepo(sql *db.SQLStorage) *ClientRepo {
	return &ClientRepo{
		SQLStorage: sql,
	}
}

func (clientRepo *ClientRepo) CreateClient(ctx context.Context, arg db.CreateClientParams) (db.ClientClient, error) {
	var result db.ClientClient

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateClient(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *ClientRepo) GetClient(ctx context.Context, id int64) (db.ClientClient, error) {
	var result db.ClientClient

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetClient(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *ClientRepo) GetClients(ctx context.Context, arg db.GetClientsParams) ([]db.ClientClient, error) {
	var result []db.ClientClient

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetClients(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *ClientRepo) GetClientsByCompanyId(ctx context.Context, arg db.GetClientsByCompanyIdParams) ([]db.ClientClient, error) {
	var result []db.ClientClient

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetClientsByCompanyId(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *ClientRepo) UpdateClient(ctx context.Context, arg db.UpdateClientParams) (db.ClientClient, error) {
	var result db.ClientClient

	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdateClient(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (clientRepo *ClientRepo) DeleteClient(ctx context.Context, id int64) error {
	err := clientRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		errDelete := q.DeleteClient(ctx, id)
		if errDelete != nil {
			return errDelete
		}

		return errDelete
	})

	return err
}
