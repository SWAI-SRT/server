package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost       string
	DBUser       string
	DBPassword   string
	DBName       string
	JWTSecret    string
	AWSRegion    string
	S3BucketName string
}

func LoadConfig() (config Config, err error) {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
		return
	}

	config = Config{
		DBHost:       os.Getenv("DB_HOST"),
		DBUser:       os.Getenv("DB_USER"),
		DBPassword:   os.Getenv("DB_PASSWORD"),
		DBName:       os.Getenv("DB_NAME"),
		JWTSecret:    os.Getenv("JWT_SECRET"),
		AWSRegion:    os.Getenv("AWS_REGION"),
		S3BucketName: os.Getenv("S3_BUCKET_NAME"),
	}

	return
}
