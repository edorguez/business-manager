package s3

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/edorguez/business-manager/services/file-svc/pkg/config"
)

const region = "us-east-1"

type S3File struct {
	FileName string
	FileData []byte
}

func UploadFiles(c *config.Config, bucketName string, folderName string, files []S3File) ([]string, error) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(c.Aws_Access_Key_Id, c.Aws_Secret_Access_Key_Id, ""),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 session: %w", err)
	}

	svc := s3.New(sess)
	var wg sync.WaitGroup
	var errUpload error
	urls := make([]string, len(files))

	for i, file := range files {
		wg.Add(1)
		go func(i int, file S3File) {
			defer wg.Done()
			_, err := svc.PutObject(&s3.PutObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(folderName + "/" + file.FileName),
				Body:   bytes.NewReader(file.FileData),
				// ACL:    aws.String("public-read"),
			})
			if err != nil {
				errUpload = fmt.Errorf("failed to upload file %s: %w", file.FileName, err)
				return
			}
			// https://business-manager-bucket-s3.s3.us-east-1.amazonaws.com/images/myImage.png
			urls[i] = fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s", bucketName, region, folderName, file.FileName)
		}(i, file)
	}

	wg.Wait()

	return urls, errUpload
}

func DeleteFiles(c *config.Config, bucketName string, folderName string, files []string) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(region),
		Credentials: credentials.NewStaticCredentials(c.Aws_Access_Key_Id, c.Aws_Secret_Access_Key_Id, ""),
	})
	if err != nil {
		return fmt.Errorf("failed to create S3 session: %w", err)
	}

	svc := s3.New(sess)
	var wg sync.WaitGroup
	var errDelete error

	for _, file := range files {
		wg.Add(1)
		go func(file string) {
			defer wg.Done()
			_, err := svc.DeleteObject(&s3.DeleteObjectInput{
				Bucket: aws.String(bucketName),
				Key:    aws.String(folderName + "/" + file),
			})
			if err != nil {
				errDelete = fmt.Errorf("failed to delete file %s: %w", file, err)
				return
			}
		}(file)
	}

	wg.Wait()

	return errDelete
}
