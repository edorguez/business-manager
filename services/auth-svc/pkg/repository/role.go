package repository

import (
	"context"

	db "github.com/edorguez/business-manager/services/auth-svc/pkg/db/sqlc"
)

type RoleRepo struct {
	SQLStorage *db.SQLStorage
}

func NewRoleRepo(sql *db.SQLStorage) *RoleRepo {
	return &RoleRepo{
		SQLStorage: sql,
	}
}

func (roleRepo *RoleRepo) GetRole(ctx context.Context, id int64) (db.GetRoleRow, error) {
	var result db.GetRoleRow

	err := roleRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetRole(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (roleRepo *RoleRepo) GetRoles(ctx context.Context) ([]db.GetRolesRow, error) {
	var result []db.GetRolesRow

	err := roleRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetRoles(ctx)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
