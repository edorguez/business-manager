package client

import (
	"context"
	"fmt"
	"net/http"
	"os"

	pb "github.com/EdoRguez/business-manager/file-svc/pkg/pb/file"
	"github.com/EdoRguez/business-manager/product-svc/pkg/config"
	"google.golang.org/grpc"
)

var fileServiceClient pb.FileServiceClient

func InitFileServiceClient(c *config.Config) error {
	fmt.Println("File Client :  InitFileServiceClient")

	appEnv := os.Getenv("ENVIRONMENT")
	if appEnv == "" {
		appEnv = "development" // Default to development if the variable is not set
	}

	var fileSvcUrl string
	if appEnv == "production" {
		fmt.Println("Running in production mode")
		fileSvcUrl = c.File_Svc_Url + ":" + c.File_Svc_Port
	} else {
		fmt.Println("Running in development mode")
		fileSvcUrl = c.Development_Url + ":" + c.File_Svc_Port
	}

	// using WithInsecure() because no SSL running
	cc, err := grpc.Dial(fileSvcUrl, grpc.WithInsecure())

	if err != nil {
		fmt.Println("Could not connect:", err)
		return err
	}

	fileServiceClient = pb.NewFileServiceClient(cc)
	return nil
}

type FileData struct {
	FileName string
	FileData []byte
}

func UploadFiles(bucketName string, folderName string, files []FileData, c context.Context) (*pb.UploadFilesResponse, error) {
	fmt.Println("File CLIENT :  UploadFiles")

	params := &pb.UploadFilesRequest{
		BucketName: bucketName,
		FolderName: folderName,
		Files:      make([]*pb.FileData, len(files)),
	}

	for i, file := range files {
		params.Files[i] = &pb.FileData{
			FileName: file.FileName,
			FileData: file.FileData,
		}
	}

	res, err := fileServiceClient.UploadFiles(c, params)

	if err != nil {
		fmt.Println("File CLIENT :  UploadFiles - ERROR")
		fmt.Println(err.Error())

		return nil, err
	}

	fmt.Println("File CLIENT :  UploadFiles - SUCCESS")

	if res.Status != http.StatusOK {
		return nil, err
	}

	return &pb.UploadFilesResponse{
		FileUrls: res.FileUrls,
		Status:   http.StatusOK,
	}, nil
}
