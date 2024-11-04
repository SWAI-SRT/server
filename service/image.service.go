package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"swai/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ImageService struct {
	s3Client *s3.Client
	bucket   string
	region   string
}

func NewImageService(cfg *config.Config) *ImageService {
	awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(), awsConfig.WithRegion(cfg.AWSRegion))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	return &ImageService{
		s3Client: s3.NewFromConfig(awsCfg),
		bucket:   cfg.S3BucketName,
		region:   cfg.AWSRegion,
	}
}

func (s *ImageService) Upload(fileName string, file multipart.File) (string, error) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(file)
	if err != nil {
		return "", err
	}

	_, err = s.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.bucket),
		Key:    aws.String(fileName),
		Body:   buffer,
	})
	if err != nil {
		return "", err
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, fileName)
	return url, nil
}
