package repository

import (
	"context"

	db "github.com/edorguez/business-manager/services/company-svc/pkg/db/sqlc"
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

func (companyRepo *CompanyRepo) GetCompany(ctx context.Context, id int64) (db.GetCompanyRow, error) {
	var result db.GetCompanyRow

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

func (companyRepo *CompanyRepo) GetCompanyByName(ctx context.Context, name string) (db.GetCompanyByNameRow, error) {
	var result db.GetCompanyByNameRow

	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCompanyByName(ctx, name)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (companyRepo *CompanyRepo) GetCompanyByNameUrl(ctx context.Context, nameUrl string) (db.GetCompanyByNameUrlRow, error) {
	var result db.GetCompanyByNameUrlRow

	err := companyRepo.SQLStorage.ExecTx(ctx, func(q *db.Queries) error {
		var err error

		result, err = q.GetCompanyByNameUrl(ctx, nameUrl)
		if err != nil {
			return err
		}

		return err
	})

	return result, err
}

func (companyRepo *CompanyRepo) GetCompanies(ctx context.Context, arg db.GetCompaniesParams) ([]db.GetCompaniesRow, error) {
	var result []db.GetCompaniesRow

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
