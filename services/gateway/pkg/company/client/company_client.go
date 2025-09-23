package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/edorguez/business-manager/services/gateway/pkg/company/contracts"
	"github.com/edorguez/business-manager/services/gateway/pkg/config"
	pb "github.com/edorguez/business-manager/shared/pb/company"
	"github.com/edorguez/business-manager/shared/types"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var companyServiceClient pb.CompanyServiceClient

func InitCompanyServiceClient(c *config.Config) error {
	fmt.Println("API Gateway :  InitCompanyServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var companySvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		companySvcUrl = c.ProductionUrl + ":" + c.CompanySvcPort
	} else {
		fmt.Println("Running in development mode")
		companySvcUrl = c.DevelopmentUrl + ":" + c.CompanySvcPort
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

func CreateCompany(body contracts.CreateCompanyRequest, c context.Context) (*pb.CreateCompanyResponse, *types.Error) {
	fmt.Println("Company CLIENT :  CreateCompany")

	fmt.Println("Company CLIENT :  CreateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createCompanyParams := &pb.CreateCompanyRequest{
		Name:            body.Name,
		LastPaymentDate: timestamppb.New(body.LastPaymentDate),
	}

	res, err := companyServiceClient.CreateCompany(c, createCompanyParams)

	if err != nil {
		fmt.Println("Company CLIENT :  CreateCompany - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  CreateCompany - SUCCESS")
	return res, nil
}

func GetCompany(id int64, c context.Context) (*contracts.GetCompanyResponse, *types.Error) {
	fmt.Println("Company CLIENT :  GetCompany")

	params := &pb.GetCompanyRequest{
		Id: id,
	}

	res, err := companyServiceClient.GetCompany(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompany - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  GetCompany - SUCCESS")

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetCompanyResponse{
		Id:              res.Id,
		Name:            res.Name,
		NameFormatUrl:   res.NameFormatUrl,
		ImageUrl:        res.ImageUrl,
		PlanId:          res.PlanId,
		LastPaymentDate: res.LastPaymentDate.AsTime(),
	}, nil
}

func GetCompanyByName(name string, c context.Context) (*contracts.GetCompanyByNameResponse, *types.Error) {
	fmt.Println("Company CLIENT :  GetCompanyByName")

	params := &pb.GetCompanyByNameRequest{
		Name: name,
	}

	res, err := companyServiceClient.GetCompanyByName(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompanyByName - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  GetCompanyByName - SUCCESS")

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetCompanyByNameResponse{
		Id:              res.Id,
		Name:            res.Name,
		NameFormatUrl:   res.NameFormatUrl,
		ImageUrl:        res.ImageUrl,
		PlanId:          res.PlanId,
		LastPaymentDate: res.LastPaymentDate.AsTime(),
	}, nil
}

func GetCompanyByNameUrl(nameUrl string, c context.Context) (*contracts.GetCompanyByNameUrlResponse, *types.Error) {
	fmt.Println("Company CLIENT :  GetCompanyByNameUrl")

	params := &pb.GetCompanyByNameUrlRequest{
		NameUrl: nameUrl,
	}

	res, err := companyServiceClient.GetCompanyByNameUrl(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompanyByNameUrl - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  GetCompanyByNameUrl - SUCCESS")

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetCompanyByNameUrlResponse{
		Id:              res.Id,
		Name:            res.Name,
		NameFormatUrl:   res.NameFormatUrl,
		ImageUrl:        res.ImageUrl,
		PlanId:          res.PlanId,
		LastPaymentDate: res.LastPaymentDate.AsTime(),
	}, nil
}

func GetCompanies(params *pb.GetCompaniesRequest, c context.Context) ([]*contracts.GetCompanyResponse, *types.Error) {
	fmt.Println("Company CLIENT :  GetCompanies")

	res, err := companyServiceClient.GetCompanies(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompanies - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &types.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	cr := make([]*contracts.GetCompanyResponse, 0, len(res.Companies))
	for _, v := range res.Companies {
		cr = append(cr, &contracts.GetCompanyResponse{
			Id:              v.Id,
			Name:            v.Name,
			NameFormatUrl:   v.NameFormatUrl,
			ImageUrl:        v.ImageUrl,
			PlanId:          v.PlanId,
			LastPaymentDate: v.LastPaymentDate.AsTime(),
		})
	}

	fmt.Println("Company CLIENT :  GetCompanies - SUCCESS")
	return cr, nil
}

func UpdateCompany(id int64, body contracts.UpdateCompanyRequest, image []byte, c context.Context) (*pb.UpdateCompanyResponse, *types.Error) {
	fmt.Println("Company CLIENT :  UpdateCompany")

	fmt.Println("Company CLIENT :  UpdateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateCompanyParams := &pb.UpdateCompanyRequest{
		Id:            int64(id),
		Name:          body.Name,
		NameFormatUrl: body.NameFormatUrl,
		Image:         image,
	}

	res, err := companyServiceClient.UpdateCompany(c, updateCompanyParams)

	if err != nil {
		fmt.Println("Company CLIENT :  UpdateCompany - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  UpdateCompany - SUCCESS")
	return res, nil
}

func DeleteCompany(id int64, c context.Context) (*pb.DeleteCompanyResponse, *types.Error) {
	fmt.Println("Company CLIENT :  DeleteCompany")

	params := &pb.DeleteCompanyRequest{
		Id: id,
	}

	res, err := companyServiceClient.DeleteCompany(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  DeleteCompany - ERROR")
		fmt.Println(err.Error())

		error := &types.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  DeleteCompany - SUCCESS")
	return res, nil
}
