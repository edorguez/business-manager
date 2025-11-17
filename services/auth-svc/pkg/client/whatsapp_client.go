package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/auth-svc/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/whatsapp"
	"google.golang.org/grpc"
)

var whatsappServiceClient pb.WhatsappServiceClient

func InitWhatsappServiceClient(c *config.Config) error {
	fmt.Println("Whatsapp Client :  InitWhatsappServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var whatsappSvcUrl string
	if appEnv == "development" {
		fmt.Println("Running in development mode")
		whatsappSvcUrl = "localhost:" + c.WhatsappSvcPort
	} else {
		fmt.Println("Running in docker mode")
		whatsappSvcUrl = c.WhatsappSvcUrl + ":" + c.WhatsappSvcPort
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

func CreateBusinessPhone(params *pb.CreateBusinessPhoneRequest, c context.Context) (*pb.CreateBusinessPhoneResponse, error) {
	fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone")

	res, err := whatsappServiceClient.CreateBusinessPhone(c, params)

	if err != nil || res.Error != "" {
		fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	if res.Status != http.StatusNoContent {
		return nil, err
	}

	fmt.Println("Whatsapp CLIENT :  CreateBusinessPhone - SUCCESS")

	return &pb.CreateBusinessPhoneResponse{
		Status: res.Status,
	}, nil
}
