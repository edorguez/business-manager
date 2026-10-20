package s3

import (
	"bytes"
	"fmt"
	"mime"
	"net/url"
	"path/filepath"
	"strings"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/edorguez/business-manager/services/file-svc/pkg/config"
)

type S3File struct {
	FileName string
	FileData []byte
}

func newS3Client(c *config.Config) (*s3.S3, error) {
	region := c.S3Region
	if region == "" {
		region = "us-east-1"
	}

	awsConfig := &aws.Config{
		Region:           aws.String(region),
		Credentials:      credentials.NewStaticCredentials(c.AwsAccessKeyId, c.AwsSecretAccessKeyId, ""),
		S3ForcePathStyle: aws.Bool(c.S3ForcePathStyle),
	}

	// VersityGW/MinIO (and other S3-compatible services) need a custom endpoint.
	if c.S3Endpoint != "" {
		awsConfig.Endpoint = aws.String(c.S3Endpoint)
	}

	sess, err := session.NewSession(awsConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 session: %w", err)
	}

	return s3.New(sess), nil
}

func publicObjectURL(c *config.Config, bucketName string, folderName string, fileName string) string {
	base := strings.TrimRight(c.S3PublicUrl, "/")
	if base == "" {
		base = "/media"
	}

	return fmt.Sprintf(
		"%s/%s/%s/%s",
		base,
		bucketName,
		strings.Trim(folderName, "/"),
		url.PathEscape(fileName),
	)
}

func contentType(fileName string) string {
	if value := mime.TypeByExtension(filepath.Ext(fileName)); value != "" {
		return value
	}
	return "application/octet-stream"
}

func UploadFiles(c *config.Config, bucketName string, folderName string, files []S3File) ([]string, error) {
	svc, err := newS3Client(c)
	if err != nil {
		return nil, err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	var errUpload error
	urls := make([]string, len(files))

	for i, file := range files {
		wg.Add(1)
		go func(i int, file S3File) {
			defer wg.Done()
			_, err := svc.PutObject(&s3.PutObjectInput{
				Bucket:      aws.String(bucketName),
				Key:         aws.String(folderName + "/" + file.FileName),
				Body:        bytes.NewReader(file.FileData),
				ContentType: aws.String(contentType(file.FileName)),
			})
			if err != nil {
				mu.Lock()
				errUpload = fmt.Errorf("failed to upload file %s: %w", file.FileName, err)
				mu.Unlock()
				return
			}
			urls[i] = publicObjectURL(c, bucketName, folderName, file.FileName)
		}(i, file)
	}

	wg.Wait()

	return urls, errUpload
}

func DeleteFiles(c *config.Config, bucketName string, folderName string, files []string) error {
	svc, err := newS3Client(c)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
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
				mu.Lock()
				errDelete = fmt.Errorf("failed to delete file %s: %w", file, err)
				mu.Unlock()
				return
			}
		}(file)
	}

	wg.Wait()

	return errDelete
}
