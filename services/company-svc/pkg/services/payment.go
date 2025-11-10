package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/edorguez/business-manager/services/company-svc/pkg/db/sqlc"
	repo "github.com/edorguez/business-manager/services/company-svc/pkg/repository"
	"github.com/edorguez/business-manager/shared/constants"
	"github.com/edorguez/business-manager/shared/pb/payment"
	"github.com/edorguez/business-manager/shared/util/type_converter"
)

type PaymentService struct {
	Repo        *repo.PaymentRepo
	CompanyRepo *repo.CompanyRepo
	payment.UnimplementedPaymentServiceServer
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	fmt.Println("Payment Service :  CreatePayment")
	fmt.Println("Payment Service :  CreatePayment - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	company, errCompany := s.CompanyRepo.GetCompany(ctx, req.CompanyId)

	if errCompany != nil {
		fmt.Println("Payment Service :  CreatePayment - ERROR")
		return &payment.CreatePaymentResponse{
			Status: http.StatusInternalServerError,
			Error:  errCompany.Error(),
		}, nil
	}

	if company.PlanID == constants.PLAN_ID_BASIC {
		ps, err := s.Repo.GetPayments(ctx, db.GetPaymentsParams{
			CompanyID: req.CompanyId,
			Offset:    0,
			Limit:     constants.MAX_BASIC_PLAN_PAYMENTS,
		})

		if err != nil {
			fmt.Println("Payment Service :  CreatePayment - ERROR")
			fmt.Println(err.Error())
			return &payment.CreatePaymentResponse{
				Status: http.StatusInternalServerError,
				Error:  err.Error(),
			}, nil
		}

		if len(ps) >= constants.MAX_BASIC_PLAN_PAYMENTS {
			fmt.Println("Payment Service :  CreatePayment - ERROR")
			return &payment.CreatePaymentResponse{
				Status: http.StatusUnauthorized,
				Error:  "Can't create payment, upgrade your plan to create more payments",
			}, nil
		}
	}

	createPaymentParams := db.CreatePaymentParams{
		CompanyID:            req.CompanyId,
		Name:                 req.Name,
		Bank:                 type_converter.NewSqlNullString(req.Bank),
		AccountNumber:        type_converter.NewSqlNullString(req.AccountNumber),
		AccountType:          type_converter.NewSqlNullString(req.AccountType),
		IdentificationNumber: type_converter.NewSqlNullString(req.IdentificationNumber),
		IdentificationType:   type_converter.NewSqlNullString(req.IdentificationType),
		Phone:                type_converter.NewSqlNullString(req.Phone),
		Email:                type_converter.NewSqlNullString(req.Email),
		PaymentTypeID:        req.PaymentTypeId,
	}

	c, err := s.Repo.CreatePayment(ctx, createPaymentParams)
	if err != nil {
		fmt.Println("Payment Service :  CreatePayment - ERROR")
		fmt.Println(err.Error())
		return &payment.CreatePaymentResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Payment Service :  CreatePayment - SUCCESS")
	return &payment.CreatePaymentResponse{
		Status: http.StatusCreated,
		Id:     c.ID,
	}, nil
}

func (s *PaymentService) GetPayment(ctx context.Context, req *payment.GetPaymentRequest) (*payment.GetPaymentResponse, error) {
	fmt.Println("Payment Service :  GetPayment")
	fmt.Println("Payment Service :  GetPayment - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	p, err := s.Repo.GetPayment(ctx, req.Id)
	if err != nil {
		fmt.Println("Payment Service :  GetPayment - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &payment.GetPaymentResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Payment Service :  GetPayment - SUCCESS")
	return &payment.GetPaymentResponse{
		Id:                   p.ID,
		CompanyId:            p.CompanyID,
		Name:                 p.Name,
		Bank:                 type_converter.NewString(p.Bank),
		AccountNumber:        type_converter.NewString(p.AccountNumber),
		AccountType:          type_converter.NewString(p.AccountType),
		IdentificationNumber: type_converter.NewString(p.IdentificationNumber),
		IdentificationType:   type_converter.NewString(p.IdentificationType),
		Phone:                type_converter.NewString(p.Phone),
		Email:                type_converter.NewString(p.Email),
		PaymentTypeId:        p.PaymentTypeID,
		IsActive:             p.IsActive,
		PaymentType: &payment.GetChildPaymentTypeResponse{
			Id:        p.CompanyPaymentType.ID,
			Name:      p.CompanyPaymentType.Name,
			ImagePath: type_converter.NewString(p.CompanyPaymentType.ImagePath),
		},
		Status: http.StatusOK,
	}, nil
}

func (s *PaymentService) GetPayments(ctx context.Context, req *payment.GetPaymentsRequest) (*payment.GetPaymentsResponse, error) {
	fmt.Println("Payment Service :  GetPayments")
	fmt.Println("Payment Service :  GetPayments - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetPaymentsParams{
		CompanyID:     req.CompanyId,
		PaymentTypeID: req.PaymentTypeId,
		Limit:         req.Limit,
		Offset:        req.Offset,
	}

	p, err := s.Repo.GetPayments(ctx, params)
	if err != nil {
		fmt.Println("Payment Service :  GetPayments - ERROR")
		fmt.Println(err.Error())
		return &payment.GetPaymentsResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	payments := make([]*payment.GetPaymentResponse, 0, len(p))
	for _, v := range p {
		payments = append(payments, &payment.GetPaymentResponse{
			Id:                   v.ID,
			CompanyId:            v.CompanyID,
			Name:                 v.Name,
			Bank:                 type_converter.NewString(v.Bank),
			AccountNumber:        type_converter.NewString(v.AccountNumber),
			AccountType:          type_converter.NewString(v.AccountType),
			IdentificationNumber: type_converter.NewString(v.IdentificationNumber),
			IdentificationType:   type_converter.NewString(v.IdentificationType),
			Phone:                type_converter.NewString(v.Phone),
			Email:                type_converter.NewString(v.Email),
			PaymentTypeId:        v.PaymentTypeID,
			IsActive:             v.IsActive,
			PaymentType: &payment.GetChildPaymentTypeResponse{
				Id:        v.CompanyPaymentType.ID,
				Name:      v.CompanyPaymentType.Name,
				ImagePath: type_converter.NewString(v.CompanyPaymentType.ImagePath),
			}, Status: http.StatusOK,
		})
	}

	fmt.Println("Payment Service :  GetPayments - SUCCESS")
	return &payment.GetPaymentsResponse{
		Payments: payments,
		Status:   http.StatusOK,
	}, nil
}

func (s *PaymentService) GetPaymentsTypes(ctx context.Context, req *payment.GetPaymentsTypesRequest) (*payment.GetPaymentsTypesResponse, error) {
	fmt.Println("Payment Service :  GetPaymentsTypes")
	fmt.Println("Payment Service :  GetPaymentsTypes - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	p, err := s.Repo.GetPaymentsTypes(ctx, req.CompanyId)
	if err != nil {
		fmt.Println("Payment Service :  GetPaymentsTypes - ERROR")
		fmt.Println(err.Error())
		return &payment.GetPaymentsTypesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	payments := make([]*payment.PaymentType, 0, len(p))
	for _, v := range p {
		payments = append(payments, &payment.PaymentType{
			Id:            v.ID,
			CompanyId:     v.CompanyID,
			PaymentTypeId: v.PaymentTypeID,
			IsActive:      v.IsActive,
			PaymentType: &payment.GetChildPaymentTypeResponse{
				Id:        v.CompanyPaymentType.ID,
				Name:      v.CompanyPaymentType.Name,
				ImagePath: type_converter.NewString(v.CompanyPaymentType.ImagePath),
			},
		})
	}

	fmt.Println("Payment Service :  GetPaymentsTypes - SUCCESS")
	return &payment.GetPaymentsTypesResponse{
		PaymentsTypes: payments,
		Status:        http.StatusOK,
	}, nil
}

func (s *PaymentService) UpdatePayment(ctx context.Context, req *payment.UpdatePaymentRequest) (*payment.UpdatePaymentResponse, error) {
	fmt.Println("Payment Service :  UpdatePayment")
	fmt.Println("Payment Service :  UpdatePayment - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdatePaymentParams{
		ID:                   req.Id,
		Name:                 req.Name,
		Bank:                 type_converter.NewSqlNullString(req.Bank),
		AccountNumber:        type_converter.NewSqlNullString(req.AccountNumber),
		AccountType:          type_converter.NewSqlNullString(req.AccountType),
		IdentificationNumber: type_converter.NewSqlNullString(req.IdentificationNumber),
		IdentificationType:   type_converter.NewSqlNullString(req.IdentificationType),
		Phone:                type_converter.NewSqlNullString(req.Phone),
		Email:                type_converter.NewSqlNullString(req.Email),
		PaymentTypeID:        req.PaymentTypeId,
	}

	_, err := s.Repo.UpdatePayment(ctx, params)
	if err != nil {
		fmt.Println("Payment Service :  UpdatePayment - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &payment.UpdatePaymentResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Payment Service :  UpdatePayment - SUCCESS")
	return &payment.UpdatePaymentResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *PaymentService) UpdatePaymentStatus(ctx context.Context, req *payment.UpdatePaymentStatusRequest) (*payment.UpdatePaymentStatusResponse, error) {
	fmt.Println("Payment Service :  UpdatePaymentStatus")
	fmt.Println("Payment Service :  UpdatePaymentStatus - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.UpdatePaymentStatusParams{
		ID:       req.Id,
		IsActive: req.Status,
	}

	_, err := s.Repo.UpdatePaymentStatus(ctx, params)
	if err != nil {
		fmt.Println("Payment Service :  UpdatePaymentStatus - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &payment.UpdatePaymentStatusResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("Payment Service :  UpdatePaymentStatus - SUCCESS")
	return &payment.UpdatePaymentStatusResponse{
		Status: http.StatusNoContent,
	}, nil
}

func (s *PaymentService) DeletePayment(ctx context.Context, req *payment.DeletePaymentRequest) (*payment.DeletePaymentResponse, error) {
	fmt.Println("Payment Service :  DeletePayment")
	fmt.Println("Payment Service :  DeletePayment - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s.Repo.DeletePayment(ctx, req.Id)
	if err != nil {
		fmt.Println("Payment Service :  DeletePayment - ERROR")
		fmt.Println(err.Error())
		return &payment.DeletePaymentResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("Payment Service :  DeletePayment - SUCCESS")
	return &payment.DeletePaymentResponse{
		Status: http.StatusNoContent,
	}, nil
}
