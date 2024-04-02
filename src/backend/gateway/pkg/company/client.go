package company

import (
	"fmt"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"google.golang.org/grpc"
)

type ServiceClient struct {
	CompanyClient pb.CompanyServiceClient
	PaymentClient pb.PaymentServiceClient
}

func InitServiceClient(c *config.Config) ServiceClient {
	fmt.Println("API Gateway :  InitCompanyServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Company_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
	}

	return ServiceClient{
		CompanyClient: pb.NewCompanyServiceClient(cc),
		PaymentClient: pb.NewPaymentServiceClient(cc),
	}

}
