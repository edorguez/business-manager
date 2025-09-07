package services

import (
	"context"
	"fmt"
	"net/http"

	"github.com/edorguez/business-manager/services/file-svc/pkg/config"
	s3 "github.com/edorguez/business-manager/services/file-svc/pkg/s3"
	file "github.com/edorguez/business-manager/shared/pb/file"
)

type FileService struct {
	Config *config.Config
	file.UnimplementedFileServiceServer
}

func (s *FileService) UploadFiles(ctx context.Context, req *file.UploadFilesRequest) (*file.UploadFilesResponse, error) {
	fmt.Println("File Service :  UploadFiles")
	fmt.Println("File Service :  UploadFiles - Req")
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
		fmt.Println("File Service :  UploadFiles - ERROR")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("File Service :  UploadFiles - SUCCESS")
	return &file.UploadFilesResponse{
		FileUrls: filesUrls,
		Status:   http.StatusOK,
	}, nil
}

func (s *FileService) DeleteFiles(ctx context.Context, req *file.DeleteFilesRequest) (*file.DeleteFilesResponse, error) {
	fmt.Println("File Service :  DeleteFiles")
	fmt.Println("File Service :  DeleteFiles - Req")
	fmt.Println(req)
	fmt.Println("----------------")

	err := s3.DeleteFiles(s.Config, req.BucketName, req.FolderName, req.FileNames)
	if err != nil {
		fmt.Println("File Service :  DeleteFiles - ERROR")
		fmt.Println(err.Error())
		return nil, err
	}

	fmt.Println("File Service :  DeleteFiles - SUCCESS")
	return &file.DeleteFilesResponse{
		Status: http.StatusOK,
	}, nil
}
