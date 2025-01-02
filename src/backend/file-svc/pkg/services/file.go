package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/EdoRguez/business-manager/file-svc/pkg/config"
	file "github.com/EdoRguez/business-manager/file-svc/pkg/pb/file"
	s3 "github.com/EdoRguez/business-manager/file-svc/pkg/s3"
)

type FileService struct {
	Config *config.Config
	file.UnimplementedFileServiceServer
}

func (s *FileService) UploadFile(ctx context.Context, req *file.UploadFilesRequest) (*file.UploadFilesResponse, error) {
	fmt.Println("File Service :  UploadFile")
	fmt.Println("File Service :  UplaodFile - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	s3Files := make([]s3.S3File, len(req.Files))
	for i, f := range req.Files {
		s3Files[i] = s3.S3File{
			FileName: f.FileName,
			FileData: f.FileData,
		}
	}

	filesUrls, err := s3.UploadFiles(s.Config, req.BucketName, req.FolderName, s3Files)
	if err != nil {
		fmt.Println("File Service :  UploadFile - ERROR")
		fmt.Println(err.Error())
		return &file.UploadFilesResponse{
			Status: http.StatusConflict,
			Error:  err.Error(),
		}, nil
	}

	fmt.Println("File Service :  UploadFile - SUCCESS")
	return &file.UploadFilesResponse{
		FileUrls: filesUrls,
		Status:   http.StatusOK,
	}, nil
}
