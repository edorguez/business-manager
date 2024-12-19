package client

import (
	"os"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/payment_type"

	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"google.golang.org/grpc"
)

var paymentTypeServiceClient pb.PaymentTypeServiceClient

func InitPaymentTypeServiceClient(c *config.Config) error {
	fmt.Println("API Gateway :  InitPaymentTypeServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var companySvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		companySvcUrl = c.Company_Svc_Url + ":" + c.Company_Svc_Port
	} else {
		fmt.Println("Running in development mode")
		companySvcUrl = c.Development_Url + ":" + c.Company_Svc_Port
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(companySvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	paymentTypeServiceClient = pb.NewPaymentTypeServiceClient(cc)
	return nil
}

func GetPaymentType(id int64, c context.Context) (*contracts.GetPaymentTypeResponse, *contracts.Error) {
	fmt.Println("PaymentType CLIENT :  GetPaymentType")

	params := &pb.GetPaymentTypeRequest{
		Id: id,
	}

	res, err := paymentTypeServiceClient.GetPaymentType(c, params)

	if err != nil {
		fmt.Println("PaymentType CLIENT :  GetPaymentType - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("PaymentType CLIENT :  GetPaymentType - SUCCESS")

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetPaymentTypeResponse{
		Id:        res.Id,
		Name:      res.Name,
		ImagePath: res.ImagePath,
	}, nil
}

func GetPaymentTypes(params *pb.GetPaymentTypesRequest, c context.Context) ([]*contracts.GetPaymentTypeResponse, *contracts.Error) {
	fmt.Println("PaymentType CLIENT :  GetPaymentTypes")

	res, err := paymentTypeServiceClient.GetPaymentTypes(c, params)

	if err != nil {
		fmt.Println("PaymentType CLIENT :  GetPaymentTypes - ERROR")
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

	pt := make([]*contracts.GetPaymentTypeResponse, 0, len(res.PaymentTypes))
	for _, v := range res.PaymentTypes {
		pt = append(pt, &contracts.GetPaymentTypeResponse{
			Id:        v.Id,
			Name:      v.Name,
			ImagePath: v.ImagePath,
		})
	}

	fmt.Println("PaymentType CLIENT :  GetPaymentTypes - SUCCESS")
	return pt, nil
}
