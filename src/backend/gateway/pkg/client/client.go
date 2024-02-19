package client

import (
	"fmt"

	"github.com/EdoRguez/business-manager/gateway/pkg/client/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.ClientServiceClient
}

func InitServiceClient(c *config.Config) pb.ClientServiceClient {
	fmt.Println("API Gateway :  InitServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Client_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewClientServiceClient(cc)
}
