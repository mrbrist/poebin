package r2

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid/v4"
)

type R2 struct {
	Client          *s3.Client
	BucketName      string
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
}

func Setup() (*R2, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r2 := &R2{
		BucketName:      os.Getenv("BUCKET_NAME"),
		AccountID:       os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		AccessKeyID:     os.Getenv("CLOUDFLARE_KEY_ID"),
		AccessKeySecret: os.Getenv("CLOUDFLARE_KEY_SECRET"),
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2.AccessKeyID, r2.AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2.AccountID))
	})

	r2.Client = client
	return r2, nil
}

func (r2 *R2) NewBuild(raw string) error {
	key := shortuuid.New()
	_, err := r2.Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      &r2.BucketName,
		Key:         aws.String(key),
		Body:        strings.NewReader(raw),
		ContentType: aws.String("text/plain"),
	})
	if err != nil {
		return err
	}
	fmt.Printf("Added build: %s\n", key)
	return nil
}

func (r2 *R2) GetBuild(key string) (string, error) {
	res, err := r2.Client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &r2.BucketName,
		Key:    aws.String(key),
	})
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(data), nil
}
