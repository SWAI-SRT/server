package service

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"mime/multipart"
	"swai/config"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsConfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type ImageService struct {
    s3Client *s3.Client
    bucket   string
    region   string
    logger   *log.Logger
}

func NewImageService(cfg *config.Config) *ImageService {
    // 로거 초기화
    logger := log.New(log.Writer(), "[ImageService] ", log.LstdFlags|log.Lmicroseconds)
    
    logger.Printf("Initializing ImageService with region: %s, bucket: %s", cfg.AWSRegion, cfg.S3BucketName)
    
    awsCfg, err := awsConfig.LoadDefaultConfig(context.TODO(),
        awsConfig.WithRegion(cfg.AWSRegion),
        awsConfig.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(
            cfg.AWSAccessKeyID,
            cfg.AWSSecretAccessKey,
            "",
        )),
    )
    if err != nil {
        logger.Fatalf("Failed to load AWS SDK config: %v", err)
    }
    
    logger.Println("AWS configuration loaded successfully")
    
    return &ImageService{
        s3Client: s3.NewFromConfig(awsCfg),
        bucket:   cfg.S3BucketName,
        region:   cfg.AWSRegion,
        logger:   logger,
    }
}

func (s *ImageService) Upload(fileName string, file multipart.File) (string, error) {
    startTime := time.Now()
    s.logger.Printf("Starting upload for file: %s", fileName)
    
    // 파일 크기 확인
    buffer := new(bytes.Buffer)
    size, err := buffer.ReadFrom(file)
    if err != nil {
        s.logger.Printf("Error reading file into buffer: %v", err)
        return "", fmt.Errorf("failed to read file: %w", err)
    }
    s.logger.Printf("File size: %d bytes", size)
    
    // S3 업로드 시도
    s.logger.Printf("Attempting S3 upload to bucket: %s with key: %s", s.bucket, fileName)
    _, err = s.s3Client.PutObject(context.TODO(), &s3.PutObjectInput{
        Bucket: aws.String(s.bucket),
        Key:    aws.String(fileName),
        Body:   buffer,
    })
    if err != nil {
        s.logger.Printf("S3 upload failed: %v", err)
        return "", fmt.Errorf("failed to upload to S3: %w", err)
    }
    
    url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", s.bucket, s.region, fileName)
    
    duration := time.Since(startTime)
    s.logger.Printf("Upload completed successfully in %v. URL: %s", duration, url)
    
    return url, nil
}