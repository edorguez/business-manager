package services

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	db "github.com/edorguez/business-manager/services/company-svc/pkg/db/sqlc"
	repo "github.com/edorguez/business-manager/services/company-svc/pkg/repository"
	pb "github.com/edorguez/business-manager/shared/pb/payment_type"
	"github.com/edorguez/business-manager/shared/util/type_converter"
)

type PaymentTypeService struct {
	Repo *repo.PaymentTypeRepo
	pb.UnimplementedPaymentTypeServiceServer
}

func (s *PaymentTypeService) GetPaymentType(ctx context.Context, req *pb.GetPaymentTypeRequest) (*pb.GetPaymentTypeResponse, error) {
	fmt.Println("PaymentType Service :  GetPaymentType")
	fmt.Println("PaymentType Service :  GetPaymentType - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	p, err := s.Repo.GetPaymentType(ctx, req.Id)
	if err != nil {
		fmt.Println("PaymentType Service :  GetPaymentType - ERROR")
		fmt.Println(err.Error())

		resErrorStatus := http.StatusConflict
		resErrorMessage := err.Error()

		if err == sql.ErrNoRows {
			resErrorStatus = http.StatusNotFound
			resErrorMessage = "Record not found"
		}

		return &pb.GetPaymentTypeResponse{
			Status: int64(resErrorStatus),
			Error:  resErrorMessage,
		}, nil
	}

	fmt.Println("PaymentType Service :  GetPaymentType - SUCCESS")
	return &pb.GetPaymentTypeResponse{
		Id:        p.ID,
		Name:      p.Name,
		ImagePath: type_converter.NewString(p.ImagePath),
		Status:    http.StatusOK,
	}, nil
}

func (s *PaymentTypeService) GetPaymentTypes(ctx context.Context, req *pb.GetPaymentTypesRequest) (*pb.GetPaymentTypesResponse, error) {
	fmt.Println("PaymentType Service :  GetPaymentTypes")
	fmt.Println("PaymentType Service :  GetPaymentTypes - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	params := db.GetPaymentTypesParams{
		Limit:  req.Limit,
		Offset: req.Offset,
	}

	p, err := s.Repo.GetPaymentTypes(ctx, params)
	if err != nil {
		fmt.Println("PaymentType Service :  GetPaymentTypes - ERROR")
		fmt.Println(err.Error())
		return &pb.GetPaymentTypesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	paymentTypes := make([]*pb.GetPaymentTypeResponse, 0, len(p))
	for _, v := range p {
		paymentTypes = append(paymentTypes, &pb.GetPaymentTypeResponse{
			Id:        v.ID,
			Name:      v.Name,
			ImagePath: type_converter.NewString(v.ImagePath),
			Status:    http.StatusOK,
		})
	}

	fmt.Println("PaymentType Service :  GetPaymentTypes - SUCCESS")
	fmt.Println(paymentTypes)
	return &pb.GetPaymentTypesResponse{
		PaymentTypes: paymentTypes,
		Status:       http.StatusOK,
	}, nil
}
