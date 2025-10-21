package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	"github.com/edorguez/business-manager/services/gateway/pkg/whatsapp/contracts"
	pb "github.com/edorguez/business-manager/shared/pb/whatsapp"
	"github.com/edorguez/business-manager/shared/types"
	"google.golang.org/grpc"
)

var whatsappServiceClient pb.WhatsappServiceClient

func InitWhatsappServiceClient(c *config.Config) error {
	fmt.Println("Whatsapp CLIENT :  InitWhatsappServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var whatsappSvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		whatsappSvcUrl = c.WhatsappSvcUrl + ":" + c.WhatsappSvcPort
	} else {
		fmt.Println("Running in development mode")
		whatsappSvcUrl = c.DevelopmentUrl + ":" + c.WhatsappSvcPort
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(whatsappSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	whatsappServiceClient = pb.NewWhatsappServiceClient(cc)
	return nil
}

func GetBusinessPhoneByCompanyId(companyId int64, c context.Context) (*contracts.GetBusinessPhoneResponse, *types.Error) {
	fmt.Println("Whatsapp CLIENT :  GetBusinessPhoneByCompanyId")

	params := &pb.GetBusinessPhoneByCompanyIdRequest{
		CompanyId: companyId,
	}

	res, err := whatsappServiceClient.GetBusinessPhoneByCompanyId(c, params)

	if err != nil {
		fmt.Println("Whatsapp CLIENT :  GetBusinessPhoneByCompanyId - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Whatsapp CLIENT :  GetBusinessPhoneByCompanyId - SUCCESS")

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetBusinessPhoneResponse{
		CompanyId: res.CompanyId,
		Phone:     res.Phone,
	}, nil
}

func CreateBusinessPhone(body contracts.CreateBusinessPhoneRequest, c context.Context) (*pb.CreateBusinessPhoneResponse, *types.Error) {
	fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone")

	fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createWhatsappParams := &pb.CreateBusinessPhoneRequest{
		CompanyId: body.CompanyId,
		Phone:     body.Phone,
	}

	res, err := whatsappServiceClient.CreateBusinessPhone(c, createWhatsappParams)

	if err != nil {
		fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone - SUCCESS")
	return res, nil
}

func UpdateBusinessPhone(body contracts.UpdateBusinessPhoneRequest, c context.Context) (*pb.UpdateBusinessPhoneResponse, *types.Error) {
	fmt.Println("Whatsapp CLIENT :  UpdateBusinessPhone")

	fmt.Println("Whatsapp CLIENT :  UpdateBusinessPhone - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateWhatsappParams := &pb.UpdateBusinessPhoneRequest{
		CompanyId: body.CompanyId,
		Phone:     body.Phone,
	}

	res, err := whatsappServiceClient.UpdateBusinessPhone(c, updateWhatsappParams)

	if err != nil {
		fmt.Println("Whatsapp CLIENT :  UpdateBusinessPhone - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Whatsapp CLIENT :  UpdateBusinessPhone - SUCCESS")
	return res, nil
}
