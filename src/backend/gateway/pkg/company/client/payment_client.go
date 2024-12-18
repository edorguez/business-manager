package client

import (
	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/payment"

	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"google.golang.org/grpc"
)

var paymentServiceClient pb.PaymentServiceClient

func InitPaymentServiceClient(c *config.Config) error {
	fmt.Println("API Gateway :  InitPaymentServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Company_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	paymentServiceClient = pb.NewPaymentServiceClient(cc)
	return nil
}

func CreatePayment(body contracts.CreatePaymentRequest, c context.Context) (*pb.CreatePaymentResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  CreatePayment")

	fmt.Println("Payment CLIENT :  CreatePayment - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createPaymentParams := &pb.CreatePaymentRequest{
		CompanyId:            body.CompanyId,
		Name:                 body.Name,
		Bank:                 body.Bank,
		AccountNumber:        body.AccountNumber,
		AccountType:          body.AccountType,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
		Phone:                body.Phone,
		Email:                body.Email,
		PaymentTypeId:        body.PaymentTypeId,
	}

	res, err := paymentServiceClient.CreatePayment(c, createPaymentParams)

	if err != nil {
		fmt.Println("Payment CLIENT :  CreatePayment - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Payment CLIENT :  CreatePayment - SUCCESS")
	return res, nil
}

func GetPayment(id int64, c context.Context) (*contracts.GetPaymentResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  GetPayment")

	params := &pb.GetPaymentRequest{
		Id: id,
	}

	res, err := paymentServiceClient.GetPayment(c, params)

	if err != nil {
		fmt.Println("Payment CLIENT :  GetPayment - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Payment CLIENT :  GetPayment - SUCCESS")

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetPaymentResponse{
		Id:                   res.Id,
		CompanyId:            res.CompanyId,
		Name:                 res.Name,
		Bank:                 res.Bank,
		AccountNumber:        res.AccountNumber,
		AccountType:          res.AccountType,
		IdentificationNumber: res.IdentificationNumber,
		IdentificationType:   res.IdentificationType,
		Phone:                res.Phone,
		Email:                res.Email,
		PaymentTypeId:        res.PaymentTypeId,
		IsActive:             res.IsActive,
		PaymentType: &contracts.GetPaymentTypeResponse{
			Id:        res.PaymentType.Id,
			Name:      res.PaymentType.Name,
			ImagePath: res.PaymentType.ImagePath,
		},
	}, nil
}

func GetPayments(params *pb.GetPaymentsRequest, c context.Context) ([]*contracts.GetPaymentResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  GetPayments")

	res, err := paymentServiceClient.GetPayments(c, params)

	if err != nil {
		fmt.Println("Payment CLIENT :  GetPayments - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	pr := make([]*contracts.GetPaymentResponse, 0, len(res.Payments))
	for _, v := range res.Payments {
		pr = append(pr, &contracts.GetPaymentResponse{
			Id:                   v.Id,
			CompanyId:            v.CompanyId,
			Name:                 v.Name,
			Bank:                 v.Bank,
			AccountNumber:        v.Bank,
			AccountType:          v.AccountType,
			IdentificationNumber: v.IdentificationNumber,
			IdentificationType:   v.IdentificationType,
			Phone:                v.Phone,
			Email:                v.Email,
			PaymentTypeId:        v.PaymentTypeId,
			IsActive:             v.IsActive,
			PaymentType: &contracts.GetPaymentTypeResponse{
				Id:        v.PaymentType.Id,
				Name:      v.PaymentType.Name,
				ImagePath: v.PaymentType.ImagePath,
			},
		})
	}

	fmt.Println("Payment CLIENT :  GetPayments - SUCCESS")
	return pr, nil
}

func GetPaymentsTypes(params *pb.GetPaymentsTypesRequest, c context.Context) ([]*contracts.GetPaymentsTypesResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  GetPaymentsTypes")

	res, err := paymentServiceClient.GetPaymentsTypes(c, params)

	if err != nil {
		fmt.Println("Payment CLIENT :  GetPaymentsTypes - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	pr := make([]*contracts.GetPaymentsTypesResponse, 0, len(res.PaymentsTypes))
	for _, v := range res.PaymentsTypes {
		pr = append(pr, &contracts.GetPaymentsTypesResponse{
			Id:            v.Id,
			CompanyId:     v.CompanyId,
			PaymentTypeId: v.PaymentTypeId,
			IsActive:      v.IsActive,
			PaymentType: &contracts.GetPaymentTypeResponse{
				Id:        v.PaymentType.Id,
				Name:      v.PaymentType.Name,
				ImagePath: v.PaymentType.ImagePath,
			},
		})
	}

	fmt.Println("Payment CLIENT :  GetPaymentsTypes - SUCCESS")
	return pr, nil
}

func UpdatePayment(id int64, body contracts.UpdatePaymentRequest, c context.Context) (*pb.UpdatePaymentResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  UpdatePayment")

	fmt.Println("Payment CLIENT :  UpdatePayment - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updatePaymentParams := &pb.UpdatePaymentRequest{
		Id:                   int64(id),
		Name:                 body.Name,
		Bank:                 body.Bank,
		AccountNumber:        body.AccountNumber,
		AccountType:          body.AccountType,
		IdentificationNumber: body.IdentificationNumber,
		IdentificationType:   body.IdentificationType,
		Phone:                body.Phone,
		Email:                body.Email,
		PaymentTypeId:        body.PaymentTypeId,
	}

	res, err := paymentServiceClient.UpdatePayment(c, updatePaymentParams)

	if err != nil {
		fmt.Println("Payment CLIENT :  UpdatePayment - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Payment CLIENT :  UpdatePayment - SUCCESS")
	return res, nil
}

func UpdatePaymentStatus(id int64, body contracts.UpdatePaymentStatusRequest, c context.Context) (*pb.UpdatePaymentStatusResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  UpdatePaymentStatus")

	fmt.Println("Payment CLIENT :  UpdatePaymentStatus - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updatePaymentParams := &pb.UpdatePaymentStatusRequest{
		Id:     int64(id),
		Status: body.Status,
	}

	res, err := paymentServiceClient.UpdatePaymentStatus(c, updatePaymentParams)

	if err != nil {
		fmt.Println("Payment CLIENT :  UpdatePaymentStatus - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Payment CLIENT :  UpdatePaymentStatus - SUCCESS")
	return res, nil
}

func DeletePayment(id int64, c context.Context) (*pb.DeletePaymentResponse, *contracts.Error) {
	fmt.Println("Payment CLIENT :  DeletePayment")

	params := &pb.DeletePaymentRequest{
		Id: id,
	}

	res, err := paymentServiceClient.DeletePayment(c, params)

	if err != nil {
		fmt.Println("Payment CLIENT :  DeletePayment - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Payment CLIENT :  DeletePayment - SUCCESS")
	return res, nil
}
