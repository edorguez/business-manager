package client

import (
	"context"
	"fmt"
	"os"

	"github.com/EdoRguez/business-manager/order-svc/pkg/config"
	pb "github.com/EdoRguez/business-manager/whatsapp-svc/pkg/pb/whatsapp"
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

func SendMessage(params *pb.SendMessageRequest, c context.Context) (*pb.SendMessageResponse, error) {
	fmt.Println("Whatsapp CLIENT :  SendMessage")

	fmt.Println("Whatsapp CLIENT :  SendMessage - Body")
	fmt.Println(params)
	fmt.Println("-----------------")

	res, err := whatsappServiceClient.SendMessage(c, params)

	if err != nil {
		fmt.Println("Whatsapp CLIENT :  SendMessage - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Whatsapp CLIENT :  SendMessage - SUCCESS")
	return res, nil
}
