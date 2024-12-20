package config

import (
	"os"
)

type Config struct {
    DBHost       string
    DBUser       string
    DBPassword   string
    DBName       string
    JWTSecret    string
    AWSRegion    string
    S3BucketName string
    AWSAccessKeyID string
    AWSSecretAccessKey string
}

func LoadConfig() (config Config, err error) {
    config = Config{
        DBHost:       os.Getenv("DB_HOST"),
        DBUser:       os.Getenv("DB_USER"),
        DBPassword:   os.Getenv("DB_PASSWORD"),
        DBName:       os.Getenv("DB_NAME"),
        JWTSecret:    os.Getenv("JWT_SECRET"),
        AWSRegion:    os.Getenv("AWS_REGION"),
        S3BucketName: os.Getenv("S3_BUCKET_NAME"),
        AWSAccessKeyID: os.Getenv("AWS_ACCESS_KEY_ID"),
        AWSSecretAccessKey: os.Getenv("AWS_SECRET_ACCESS_KEY"),
    }

    return
}