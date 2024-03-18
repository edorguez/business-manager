package repository

import (
	"context"

	db "github.com/EdoRguez/business-manager/company-svc/pkg/db/sqlc"
)

type CompanyRepo struct {
	SQLStorage *db.SQLStorage
}

func NewCompanyRepo(sql *db.SQLStorage) *CompanyRepo {
	return &CompanyRepo{
		SQLStorage: sql,
	}
}

func (companyRepo *CompanyRepo) CreateCompany(ctx context.Context, arg db.CreateCompanyParams) (db.CompanyCompany, error) {
	var result db.CompanyCompany

	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.CreateCompany(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (companyRepo *CompanyRepo) GetCompany(ctx context.Context, id int64) (db.CompanyCompany, error) {
	var result db.CompanyCompany

	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCompany(ctx, id)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (companyRepo *CompanyRepo) GetCompanies(ctx context.Context, arg db.GetCompaniesParams) ([]db.CompanyCompany, error) {
	var result []db.CompanyCompany

	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCompanies(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (companyRepo *CompanyRepo) UpdateCompany(ctx context.Context, arg db.UpdateCompanyParams) (db.CompanyCompany, error) {
	var result db.CompanyCompany

	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.UpdateCompany(ctx, arg)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (companyRepo *CompanyRepo) DeleteCompany(ctx context.Context, id int64) error {
	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		errDelete := q.DeleteCompany(ctx, id)
		if errDelete != nil {
			return errDelete
		}

		return errDelete
	})

	return err
}
