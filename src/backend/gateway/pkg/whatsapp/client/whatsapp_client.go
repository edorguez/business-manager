package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	pb "github.com/EdoRguez/business-manager/gateway/pkg/pb/whatsapp"
	"github.com/EdoRguez/business-manager/gateway/pkg/whatsapp/contracts"
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
		whatsappSvcUrl = c.Whatsapp_Svc_Url + ":" + c.Whatsapp_Svc_Port
	} else {
		fmt.Println("Running in development mode")
		whatsappSvcUrl = c.Development_Url + ":" + c.Whatsapp_Svc_Port
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

func UpdateBusinessPhone(body contracts.UpdateBusinessPhoneRequest, c context.Context) (*pb.UpdateBusinessPhoneResponse, *contracts.Error) {
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

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Whatsapp CLIENT :  UpdateBusinessPhone - SUCCESS")
	return res, nil
}
