package main

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

type cfg struct {
	BucketName      string
	AccountID       string
	AccessKeyID     string
	AccessKeySecret string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envcfg := &cfg{
		BucketName:      os.Getenv("BUCKET_NAME"),
		AccountID:       os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		AccessKeyID:     os.Getenv("CLOUDFLARE_KEY_ID"),
		AccessKeySecret: os.Getenv("CLOUDFLARE_KEY_SECRET"),
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(envcfg.AccessKeyID, envcfg.AccessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		log.Fatal(err)
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", envcfg.AccountID))
	})

	// err = CreateBuild(client, envcfg, "bobbobbobobobobobobobobobob")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	b, err := GetBuild(client, envcfg, "mvz44hssVFeAwspmxNUpxo")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(b)
}

func CreateBuild(client *s3.Client, envcfg *cfg, raw string) error {
	key := shortuuid.New()
	_, err := client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      &envcfg.BucketName,
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

func GetBuild(client *s3.Client, envcfg *cfg, key string) (string, error) {
	res, err := client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &envcfg.BucketName,
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
