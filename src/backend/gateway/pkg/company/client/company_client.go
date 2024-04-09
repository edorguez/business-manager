package client

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/gateway/pkg/company/contracts"
	"github.com/EdoRguez/business-manager/gateway/pkg/company/pb"
	"github.com/EdoRguez/business-manager/gateway/pkg/config"
	"google.golang.org/grpc"
)

var companyServiceClient pb.CompanyServiceClient

func InitCompanyServiceClient(c *config.Config) error {
	fmt.Println("API Gateway :  InitCompanyServiceClient")
	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(c.Company_Svc_Url, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	companyServiceClient = pb.NewCompanyServiceClient(cc)
	return nil
}

func CreateCompany(body contracts.CreateCompanyRequest, c context.Context) (*pb.CreateCompanyResponse, *contracts.Error) {
	fmt.Println("Company CLIENT :  CreateCompany")

	fmt.Println("Company CLIENT :  CreateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	createCompanyParams := &pb.CreateCompanyRequest{
		Name:     body.Name,
		ImageUrl: body.ImageUrl,
	}

	res, err := companyServiceClient.CreateCompany(c, createCompanyParams)

	if err != nil {
		fmt.Println("Company CLIENT :  CreateCompany - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  CreateCompany - SUCCESS")
	return res, nil
}

func GetCompany(id int64, c context.Context) (*contracts.GetCompanyResponse, *contracts.Error) {
	fmt.Println("Company CLIENT :  GetCompany")

	params := &pb.GetCompanyRequest{
		Id: id,
	}

	res, err := companyServiceClient.GetCompany(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompany - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  GetCompany - SUCCESS")

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}
		return nil, error
	}

	return &contracts.GetCompanyResponse{
		Id:       res.Id,
		Name:     res.Name,
		ImageUrl: res.ImageUrl,
	}, nil
}

func GetCompanies(params *pb.GetCompaniesRequest, c context.Context) ([]*contracts.GetCompanyResponse, *contracts.Error) {
	fmt.Println("Company CLIENT :  GetCompanies")

	res, err := companyServiceClient.GetCompanies(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  GetCompanies - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}
		return nil, error
	}

	if res.Status != http.StatusOK {
		error := &contracts.Error{
			Status: res.Status,
			Error:  res.Error,
		}

		return nil, error
	}

	cr := make([]*contracts.GetCompanyResponse, len(res.Companies))
	for _, v := range res.Companies {
		cr = append(cr, &contracts.GetCompanyResponse{
			Id:       v.Id,
			Name:     v.Name,
			ImageUrl: v.ImageUrl,
		})
	}

	fmt.Println("Company CLIENT :  GetCompanies - SUCCESS")
	return cr, nil
}

func UpdateCompany(id int64, body contracts.UpdateCompanyRequest, c context.Context) (*pb.UpdateCompanyResponse, *contracts.Error) {
	fmt.Println("Company CLIENT :  UpdateCompany")

	fmt.Println("Company CLIENT :  UpdateCompany - Body")
	fmt.Println(body)
	fmt.Println("-----------------")

	updateCompanyParams := &pb.UpdateCompanyRequest{
		Id:       int64(id),
		Name:     body.Name,
		ImageUrl: body.ImageUrl,
	}

	res, err := companyServiceClient.UpdateCompany(c, updateCompanyParams)

	if err != nil {
		fmt.Println("Company CLIENT :  UpdateCompany - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  UpdateCompany - SUCCESS")
	return res, nil
}

func DeleteCompany(id int64, c context.Context) (*pb.DeleteCompanyResponse, *contracts.Error) {
	fmt.Println("Company CLIENT :  DeleteCompany")

	params := &pb.DeleteCompanyRequest{
		Id: id,
	}

	res, err := companyServiceClient.DeleteCompany(c, params)

	if err != nil {
		fmt.Println("Company CLIENT :  DeleteCompany - ERROR")
		fmt.Println(err.Error())

		error := &contracts.Error{
			Status: http.StatusInternalServerError,
			Error:  err.Error(),
		}

		return nil, error
	}

	fmt.Println("Company CLIENT :  DeleteCompany - SUCCESS")
	return res, nil
}
