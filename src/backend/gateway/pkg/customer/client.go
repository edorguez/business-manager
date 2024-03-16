package customer

import (
	"fmt"

	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"github.com/EdoRguez/business-manager/gateway/pkg/customer/pb"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	Client pb.CustomerServiceClient
}

func InitServiceClient(c *config.Config) pb.CustomerServiceClient {
	fmt.Println("API Gateway :  InitServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Client_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return pb.NewCustomerServiceClient(cc)
}
