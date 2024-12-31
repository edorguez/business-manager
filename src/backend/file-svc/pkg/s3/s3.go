package s3

import (
	"fmt"

	"github.com/EdoRguez/business-manager/file-svc/pkg/config"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const region = "us-east-1"

type S3File struct {
	BucketName string
	FolderName string
	FileName   string
	FileData   []byte
}

func UploadFiles(c *config.Config, files []S3File) error {
	_, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
	})
	if err != nil {
		return fmt.Errorf("failed to create S3 session: %w", err)
	}

	return nil
}
