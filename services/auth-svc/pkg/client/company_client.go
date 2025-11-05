package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/auth-svc/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/company"
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
		companySvcUrl = c.CompanySvcUrl + ":" + c.CompanySvcPort
	} else {
		fmt.Println("Running in development mode")
		companySvcUrl = "localhost:" + c.CompanySvcPort
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
		NameFormatUrl:   res.NameFormatUrl,
		ImageUrl:        res.ImageUrl,
		PlanId:          res.PlanId,
		LastPaymentDate: res.LastPaymentDate,
	}, nil
}

func CreateCompany(params *pb.CreateCompanyRequest, c context.Context) (*pb.CreateCompanyResponse, error) {
	fmt.Println("Company CLIENT :  CreateCompany")

	res, err := companyServiceClient.CreateCompany(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  CreateCompany - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Company CLIENT :  CreateCompany - SUCCESS")

	if res.Status != http.StatusOK {
		return nil, err
	}

	return &pb.CreateCompanyResponse{
		Id:     res.Id,
		Status: res.Status,
	}, nil
}

func UpdateCompanyImageUrl(params *pb.UpdateCompanyImageUrlRequest, c context.Context) (*pb.UpdateCompanyImageUrlResponse, error) {
	fmt.Println("Company CLIENT :  UpdateCompanyImageUrl")

	res, err := companyServiceClient.UpdateCompanyImageUrl(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  UpdateCompanyImageUrl - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("Company CLIENT :  UpdateCompanyImageUrl - SUCCESS")

	if res.Status != http.StatusOK {
		return nil, err
	}

	return &pb.UpdateCompanyImageUrlResponse{
		Status: res.Status,
		Error:  res.Error,
	}, nil
}
