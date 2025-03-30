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
		whatsappSvcUrl = "localhost:" + c.Whatsapp_Svc_Port
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

func SendOrderCustomerMessage(params *pb.SendOrderCustomerMessageRequest, c context.Context) (*pb.SendOrderCustomerMessageResponse, error) {
	fmt.Println("Whatsapp CLIENT :  SendOrderCustomerMessage")
	fmt.Println("Whatsapp CLIENT :  SendOrderCustomerMessage - Body")
	fmt.Println(params)
	fmt.Println("-----------------")

	res, err := whatsappServiceClient.SendOrderCustomerMessage(c, params)

	if err != nil {
		fmt.Println("Whatsapp CLIENT :  SendOrderCustomerMessage - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Whatsapp CLIENT :  SendOrderCustomerMessage - SUCCESS")
	return res, nil
}

func SendOrderBusinessMessage(params *pb.SendOrderBusinessMessageRequest, c context.Context) (*pb.SendOrderBusinessMessageResponse, error) {
	fmt.Println("Whatsapp CLIENT :  SendOrderBusinessMessage")
	fmt.Println("Whatsapp CLIENT :  SendOrderBusinessMessage - Body")
	fmt.Println(params)
	fmt.Println("-----------------")

	res, err := whatsappServiceClient.SendOrderBusinessMessage(c, params)

	if err != nil {
		fmt.Println("Whatsapp CLIENT :  SendOrderBusinessMessage - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Whatsapp CLIENT :  SendOrderBusinessMessage - SUCCESS")
	return res, nil
}
