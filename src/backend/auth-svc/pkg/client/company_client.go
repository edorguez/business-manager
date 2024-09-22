package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/auth-svc/pkg/config"
	"github.com/EdoRguez/business-manager/company-svc/pkg/pb"
	"google.golang.org/grpc"
)

var companyServiceClient pb.CompanyServiceClient

func InitCompanyServiceClient(c *config.Config) error {
	fmt.Println("Company Client :  InitCompanyServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Company_Svc_Url, grpc.WithInsecure())

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
		Id:       res.Id,
		Name:     res.Name,
		ImageUrl: res.ImageUrl,
		PlanId:   res.PlanId,
	}, nil
}
