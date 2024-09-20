package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/auth-svc/pkg/db/sqlc"
)

type UserRepo struct {
	SQLStorage *db.SQLStorage
}

func NewUserRepo(sql *db.SQLStorage) *UserRepo {
	return &UserRepo{
		SQLStorage: sql,
	}
}

func (userRepo *UserRepo) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.AuthUser, error) {
	var result db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateUser(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (userRepo *UserRepo) GetUser(ctx context.Context, id int64) (db.AuthUser, error) {
	var result db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetUser(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (userRepo *UserRepo) GetUserByEmail(ctx context.Context, email string) (db.AuthUser, error) {
	var result db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetUserByEmail(ctx, email)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (userRepo *UserRepo) GetUsers(ctx context.Context, arg db.GetUsersParams) ([]db.AuthUser, error) {
	var result []db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetUsers(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (userRepo *UserRepo) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.AuthUser, error) {
	var result db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdateUser(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (userRepo *UserRepo) DeleteUser(ctx context.Context, id int64) error {
	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		errDelete := q.DeleteUser(ctx, id)
		if errDelete != nil {
			return errDelete
		}

		return errDelete
	})

	return err
}

func (userRepo *UserRepo) UpdateEmail(ctx context.Context, arg db.UpdateEmailParams) (db.AuthUser, error) {
	var result db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdateEmail(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (userRepo *UserRepo) UpdatePassword(ctx context.Context, arg db.UpdatePasswordParams) (db.AuthUser, error) {
	var result db.AuthUser

	err := userRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdatePassword(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}
