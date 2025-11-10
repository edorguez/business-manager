package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/edorguez/business-manager/services/company-svc/pkg/client"
	"github.com/edorguez/business-manager/services/company-svc/pkg/config"
	db "github.com/edorguez/business-manager/services/company-svc/pkg/db/sqlc"
	repo "github.com/edorguez/business-manager/services/company-svc/pkg/repository"
	"github.com/edorguez/business-manager/shared/constants"
	"github.com/edorguez/business-manager/shared/pb/company"
	"github.com/edorguez/business-manager/shared/util/type_converter"
	"github.com/google/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type CompanyService struct {
	Repo   *repo.CompanyRepo
	Config *config.Config
	company.UnimplementedCompanyServiceServer
}

func (s *CompanyService) CreateCompany(ctx context.Context, req *company.CreateCompanyRequest) (*company.CreateCompanyResponse, error) {
	fmt.Println("Company Service :  CreateCompany")
	fmt.Println("Company Service :  CreateCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	createCompanyParams := db.CreateCompanyParams{
		Name:            req.Name,
		NameFormatUrl:   req.NameFormatUrl,
		IsFreeTrial:     true,
		PlanID:          constants.PLAN_ID_BASIC,
		LastPaymentDate: time.Now().UTC(),
	}

	c, err := s.Repo.CreateCompany(ctx, createCompanyParams)
	if err != nil {
		fmt.Println("Company Service :  CreateCompany - ERROR")
		fmt.Println(err.Error())
		return &company.CreateCompanyResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Company Service :  CreateCompany - SUCCESS")
	return &company.CreateCompanyResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *CompanyService) GetCompany(ctx context.Context, req *company.GetCompanyRequest) (*company.GetCompanyResponse, error) {
	fmt.Println("Company Service :  GetCompany")
	fmt.Println("Company Service :  GetCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetCompany(ctx, req.Id)
	if err != nil {
		fmt.Println("Company Service :  GetCompany - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.GetCompanyResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  GetCompany - SUCCESS")
	return &company.GetCompanyResponse{
		Id:              c.ID,
		Name:            c.Name,
		NameFormatUrl:   c.NameFormatUrl,
		ImageUrl:        type_converter.NewString(c.ImageUrl),
		PlanId:          c.PlanID,
		LastPaymentDate: timestamppb.New(c.LastPaymentDate),
		Status:          http.StatusOK,
	}, nil
}

func (s *CompanyService) GetCompanyByName(ctx context.Context, req *company.GetCompanyByNameRequest) (*company.GetCompanyByNameResponse, error) {
	fmt.Println("Company Service :  GetCompanyByName")
	fmt.Println("Company Service :  GetCompanyByName - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetCompanyByName(ctx, req.Name)
	if err != nil {
		fmt.Println("Company Service :  GetCompanyByName - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.GetCompanyByNameResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  GetCompanyByName - SUCCESS")
	return &company.GetCompanyByNameResponse{
		Id:              c.ID,
		Name:            c.Name,
		NameFormatUrl:   c.NameFormatUrl,
		ImageUrl:        type_converter.NewString(c.ImageUrl),
		PlanId:          c.PlanID,
		LastPaymentDate: timestamppb.New(c.LastPaymentDate),
		Status:          http.StatusOK,
	}, nil
}

func (s *CompanyService) GetCompanyByNameUrl(ctx context.Context, req *company.GetCompanyByNameUrlRequest) (*company.GetCompanyByNameUrlResponse, error) {
	fmt.Println("Company Service :  GetCompanyByNameUrl")
	fmt.Println("Company Service :  GetCompanyByNameUrl - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	c, err := s.Repo.GetCompanyByNameUrl(ctx, req.NameUrl)
	if err != nil {
		fmt.Println("Company Service :  GetCompanyByNameUrl - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.GetCompanyByNameUrlResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  GetCompanyByNameUrl - SUCCESS")
	return &company.GetCompanyByNameUrlResponse{
		Id:              c.ID,
		Name:            c.Name,
		NameFormatUrl:   c.NameFormatUrl,
		ImageUrl:        type_converter.NewString(c.ImageUrl),
		PlanId:          c.PlanID,
		LastPaymentDate: timestamppb.New(c.LastPaymentDate),
		Status:          http.StatusOK,
	}, nil
}

func (s *CompanyService) GetCompanies(ctx context.Context, req *company.GetCompaniesRequest) (*company.GetCompaniesResponse, error) {
	fmt.Println("Company Service :  GetCompanies")
	fmt.Println("Company Service :  GetCompanies - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetCompaniesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	c, err := s.Repo.GetCompanies(ctx, params)
	if err != nil {
		fmt.Println("Company Service :  GetCompanies - ERROR")
		fmt.Println(err.Error())
		return &company.GetCompaniesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	companies := make([]*company.GetCompanyResponse, 0, len(c))
	for _, v := range c {
		companies = append(companies, &company.GetCompanyResponse{
			Id:              v.ID,
			Name:            v.Name,
			NameFormatUrl:   v.NameFormatUrl,
			ImageUrl:        type_converter.NewString(v.ImageUrl),
			PlanId:          v.PlanID,
			LastPaymentDate: timestamppb.New(v.LastPaymentDate),
			Status:          http.StatusOK,
		})
	}

	fmt.Println("Company Service :  GetCompanies - SUCCESS")
	return &company.GetCompaniesResponse{
		Companies: companies,
		Status:    http.StatusOK,
	}, nil
}

func (s *CompanyService) UpdateCompany(ctx context.Context, req *company.UpdateCompanyRequest) (*company.UpdateCompanyResponse, error) {
	fmt.Println("Company Service :  UpdateCompany")
	fmt.Println("Company Service :  UpdateCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	if err := client.InitFileServiceClient(s.Config); err != nil {
		return &company.UpdateCompanyResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	c, err := s.Repo.GetCompany(ctx, req.Id)
	if err != nil {
		fmt.Println("Company Service :  UpdateCompany - ERROR")
		fmt.Println(err.Error())

		return &company.UpdateCompanyResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	if c.ImageUrl.Valid {
		lastIndex := strings.LastIndex(c.ImageUrl.String, "/")
		imageName := c.ImageUrl.String[lastIndex+1:]

		fmt.Println("Company Service :  UpdateCompany - Image To Delete")
		fmt.Println(imageName)

		imagesToDelete := []string{imageName}
		_, err = client.DeleteFiles("business-manager-bucket-s3", "images/companies", imagesToDelete, ctx)
		if err != nil {
			fmt.Println("Company Service :  UpdateCompany - ERROR")
			fmt.Println(err.Error())
			return &company.UpdateCompanyResponse{
				Status: http.StatusConflict,
				Error:  err.Error(),
			}, nil
		}
	}

	var imageUrl string
	if req.Image != nil {
		fileData := client.FileData{
			FileName: fmt.Sprintf("company-%d-image-%s", req.Id, uuid.New()),
			FileData: req.Image,
		}

		filesToUpload := []client.FileData{fileData}

		imagesUrl, err := client.UploadFiles("business-manager-bucket-s3", "images/companies", filesToUpload, ctx)
		if err != nil {
			fmt.Println("Product Service :  UpdateCompany - ERROR")
			fmt.Println(err.Error())
			return &company.UpdateCompanyResponse{
				Status: http.StatusConflict,
				Error:  err.Error(),
			}, nil
		}

		if imagesUrl != nil && len(imagesUrl.FileUrls) > 0 {
			imageUrl = imagesUrl.FileUrls[0]
		}
	}

	params := db.UpdateCompanyParams{
		ID:            req.Id,
		Name:          req.Name,
		NameFormatUrl: req.NameFormatUrl,
		ImageUrl:      type_converter.NewSqlNullString(&imageUrl),
	}

	_, err = s.Repo.UpdateCompany(ctx, params)
	if err != nil {
		fmt.Println("Company Service :  UpdateCompany - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.UpdateCompanyResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  UpdateCompany - SUCCESS")
	return &company.UpdateCompanyResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *CompanyService) DeleteCompany(ctx context.Context, req *company.DeleteCompanyRequest) (*company.DeleteCompanyResponse, error) {
	fmt.Println("Company Service :  DeleteCompany")
	fmt.Println("Company Service :  DeleteCompany - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeleteCompany(ctx, req.Id)
	if err != nil {
		fmt.Println("Company Service :  DeleteCompany - ERROR")
		fmt.Println(err.Error())
		return &company.DeleteCompanyResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Company Service :  DeleteCompany - SUCCESS")
	return &company.DeleteCompanyResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *CompanyService) UpdateCompanyImageUrl(ctx context.Context, req *company.UpdateCompanyImageUrlRequest) (*company.UpdateCompanyImageUrlResponse, error) {
	fmt.Println("Company Service :  UpdateCompanyImageUrl")
	fmt.Println("Company Service :  UpdateCompanyImageUrl - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	if err := client.InitFileServiceClient(s.Config); err != nil {
		return &company.UpdateCompanyImageUrlResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	c, err := s.Repo.GetCompany(ctx, req.Id)
	if err != nil {
		fmt.Println("Company Service :  UpdateCompanyImageUrl - ERROR")
		fmt.Println(err.Error())

		return &company.UpdateCompanyImageUrlResponse{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}, nil
	}

	params := db.UpdateCompanyImageUrlParams{
		ID:       c.ID,
		ImageUrl: type_converter.NewSqlNullString(&req.ImageUrl),
	}

	_, err = s.Repo.UpdateCompanyImageUrl(ctx, params)
	if err != nil {
		fmt.Println("Company Service :  UpdateCompanyImageUrl - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &company.UpdateCompanyImageUrlResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Company Service :  UpdateCompanyImageUrl - SUCCESS")
	return &company.UpdateCompanyImageUrlResponse{
		Status: http.StatusNoContent,
	}, nil
}
