package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/EdoRguez/business-manager/company-svc/pkg/db/sqlc"
	payment "github.com/EdoRguez/business-manager/company-svc/pkg/pb"
	repo "github.com/EdoRguez/business-manager/company-svc/pkg/repository"
	"github.com/EdoRguez/business-manager/company-svc/pkg/util/type_converter"
)

type PaymentService struct {
	Repo *repo.PaymentRepo
	payment.UnimplementedPaymentServiceServer
}

func (s *PaymentService) CreatePayment(ctx context.Context, req *payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	fmt.Println("Payment Service :  CreatePayment")
	fmt.Println("Payment Service :  CreatePayment - Req")
	fmt.Println(req)
	fmt.Println("----------------")

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
		Bank:                 p.Bank.String,
		AccountNumber:        p.AccountNumber.String,
		AccountType:          p.AccountType.String,
		IdentificationNumber: p.IdentificationNumber.String,
		IdentificationType:   p.IdentificationType.String,
		Phone:                p.Phone.String,
		Email:                p.Email.String,
		PaymentTypeId:        p.PaymentTypeID,
		PaymentType: &payment.GetPaymentTypeResponse{
			Id:   p.CompanyPaymentType.ID,
			Name: p.CompanyPaymentType.Name,
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
		CompanyID: req.CompanyId,
		Limit:     req.Limit,
		Offset:    req.Offset,
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

	var payments []*payment.GetPaymentResponse
	for _, v := range p {
		payments = append(payments, &payment.GetPaymentResponse{
			Id:                   v.ID,
			CompanyId:            v.CompanyID,
			Name:                 v.Name,
			Bank:                 v.Bank.String,
			AccountNumber:        v.AccountNumber.String,
			AccountType:          v.AccountType.String,
			IdentificationNumber: v.IdentificationNumber.String,
			IdentificationType:   v.IdentificationType.String,
			Phone:                v.Phone.String,
			Email:                v.Email.String,
			PaymentTypeId:        v.PaymentTypeID,
			PaymentType: &payment.GetPaymentTypeResponse{
				Id:   v.CompanyPaymentType.ID,
				Name: v.CompanyPaymentType.Name,
			}, Status: http.StatusOK,
		})
	}

	fmt.Println("Payment Service :  GetPayments - SUCCESS")
	return &payment.GetPaymentsResponse{
		Payments: payments,
		Status:   http.StatusOK,
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
