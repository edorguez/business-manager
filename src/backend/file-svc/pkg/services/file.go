package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/file-svc/pkg/config"
	file "github.com/EdoRguez/business-manager/file-svc/pkg/pb/file"
)

type FileService struct {
	Config *config.Config
	file.UnimplementedFileServiceServer
}

func (s *FileService) UploadFile(ctx context.Context, req *file.UploadFileRequest) (*file.UploadFileResponse, error) {
	fmt.Println("File Service :  UploadFile")
	fmt.Println("File Service :  UplaodFile - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	errChan := make(chan error)
	// go func() {
	// 	defer close(errChan)
	// 	_, err := client.CreateCustomer(&customer.CreateCustomerRequest{
	// 		CompanyId:            req.CompanyId,
	// 		FirstName:            req.Customer.FirstName,
	// 		LastName:             req.Customer.LastName,
	// 		Phone:                &req.Customer.Phone,
	// 		IdentificationNumber: req.Customer.IdentificationNumber,
	// 		IdentificationType:   req.Customer.IdentificationType,
	// 	}, ctx)

	// 	if err != nil {
	// 		fmt.Println("File Service :  UploadFile - ERROR")
	// 		fmt.Println(err.Error())
	// 		errChan <- err
	// 		return
	// 	}

	// 	errChan <- nil
	// }()

	if errS3 := <-errChan; errS3 != nil {
		return &file.UploadFileResponse{
			Status: http.StatusInternalServerError,
			Error:  errS3.Error(),
		}, nil
	}

	fmt.Println("File Service :  UploadFile - SUCCESS")
	return &file.UploadFileResponse{
		Status: http.StatusCreated,
	}, nil
}
