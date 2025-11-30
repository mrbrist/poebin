package r2

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"github.com/lithammer/shortuuid/v4"
	"github.com/mrbrist/poebin/internal/utils"
)

type r2 struct {
	client          *s3.Client
	bucketName      string
	accountID       string
	accessKeyID     string
	accessKeySecret string
}

type Build struct {
	LastModified *time.Time
	Id           string
	Raw          string
	Data         *utils.PathOfBuilding
}

func Setup() (*r2, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	r2 := &r2{
		bucketName:      os.Getenv("BUCKET_NAME"),
		accountID:       os.Getenv("CLOUDFLARE_ACCOUNT_ID"),
		accessKeyID:     os.Getenv("CLOUDFLARE_KEY_ID"),
		accessKeySecret: os.Getenv("CLOUDFLARE_KEY_SECRET"),
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(r2.accessKeyID, r2.accessKeySecret, "")),
		config.WithRegion("auto"),
	)
	if err != nil {
		return nil, err
	}

	client := s3.NewFromConfig(cfg, func(o *s3.Options) {
		o.BaseEndpoint = aws.String(fmt.Sprintf("https://%s.r2.cloudflarestorage.com", r2.accountID))
	})

	r2.client = client
	return r2, nil
}

func (r2 *r2) NewBuild(raw string) (string, error) {
	id := shortuuid.New()
	_, err := r2.client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket:      &r2.bucketName,
		Key:         aws.String(id),
		Body:        strings.NewReader(raw),
		ContentType: aws.String("text/plain"),
	})
	if err != nil {
		return "", err
	}
	return id, nil
}

func (r2 *r2) GetBuild(id string) (*Build, error) {
	res, err := r2.client.GetObject(context.Background(), &s3.GetObjectInput{
		Bucket: &r2.bucketName,
		Key:    aws.String(id),
	})
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// Construct build struct for use in template, decode pob export here
	decoded, err := utils.RawToGo(string(data))
	if err != nil {
		return nil, err
	}
	build := Build{
		LastModified: res.LastModified,
		Id:           id,
		Raw:          string(data),
		Data:         decoded,
	}

	return &build, nil
}
