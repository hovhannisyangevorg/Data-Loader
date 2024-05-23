package AWSClient

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/config/DataProcesorConfig"
	"github.com/hovhannisyangevorg/Data-Procesor/internal/utils"
	"strings"
)

type AWSServiceMethod interface {
	CreateAWSService(cfg *DataProcesorConfig.Config) error
	CreateAWSSession(cfg *DataProcesorConfig.Config) error
	CreateBucket(cfg *DataProcesorConfig.Config) error
}

type AWSService struct {
	AWSSession *session.Session
	S3service  *s3.S3
}

func NewAWSService(cfg *DataProcesorConfig.Config) (*AWSService, error) {
	service := &AWSService{}
	err := service.CreateAWSService(cfg)
	if err != nil {
		return nil, utils.WrapError("NewAWSService", err)
	}
	err = service.CreateBucket(cfg)
	if err != nil {
		return nil, utils.WrapError("NewAWSService", err)
	}
	return service, nil
}

func (s *AWSService) CreateAWSService(cfg *DataProcesorConfig.Config) error {
	err := s.CreateAWSSession(cfg)
	if err != nil {
		return utils.WrapError("CreateAWSService", err)
	}
	service := s3.New(s.AWSSession)
	s.S3service = service
	return nil
}

func (s *AWSService) CreateAWSSession(cfg *DataProcesorConfig.Config) error {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(cfg.GetAwsRegion()),
		Credentials: credentials.NewStaticCredentials(cfg.GetAwsKye(), cfg.GetAwsSecret(), ""),
	})
	if err != nil {
		return utils.WrapError("CreateAWSSession", err)
	}
	s.AWSSession = sess
	return nil
}

func (s *AWSService) CreateBucket(cfg *DataProcesorConfig.Config) error {
	_, err := s.S3service.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(cfg.GetAwsS3BucketName()),
	})
	if err != nil {
		return utils.WrapError("CreateBucket", err)
	}
	return nil
}

func (s *AWSService) UploadFile(cfg *DataProcesorConfig.Config, fileName string, fileContent []byte) error {
	_, err := s.S3service.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(cfg.GetAwsS3BucketName()),
		Key:    aws.String(fileName),
		Body:   aws.ReadSeekCloser(strings.NewReader(string(fileContent))),
	})
	if err != nil {
		return utils.WrapError("UploadFile", err)
	}
	fmt.Printf("Successfully uploaded %q to bucket %q\n", fileName, cfg.GetAwsS3BucketName())
	return nil
}
