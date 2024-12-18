package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/EdoRguez/business-manager/auth-svc/pkg/config"
	pb "github.com/EdoRguez/business-manager/company-svc/pkg/pb/company"
	"google.golang.org/grpc"
)

var companyServiceClient pb.CompanyServiceClient

func InitCompanyServiceClient(c *config.Config) error {
	fmt.Println("Company Client :  InitCompanyServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var companySvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		companySvcUrl = c.Production_Url + ":" + c.Company_Svc_Port
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

	companyServiceClient = pb.NewCompanyServiceClient(cc)
	return nil
}

func GetCompany(id int64, c context.Context) (*pb.GetCompanyResponse, error) {
	fmt.Println("Company CLIENT :  GetCompany")

	params := &pb.GetCompanyRequest{
		Id: id,
	}

	res, err := companyServiceClient.GetCompany(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompany - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Company CLIENT :  GetCompany - SUCCESS")

	if res.Status != http.StatusOK {
		return nil, err
	}

	return &pb.GetCompanyResponse{
		Id:              res.Id,
		Name:            res.Name,
		ImageUrl:        res.ImageUrl,
		PlanId:          res.PlanId,
		LastPaymentDate: res.LastPaymentDate,
	}, nil
}
