package document

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

func NewS3Storage(bucket string) (*S3Storage, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("unable to load AWS config: %w", err)
	}

	client := s3.NewFromConfig(cfg)

	return &S3Storage{
		Client:     client,
		BucketName: bucket,
	}, nil
}

type S3Storage struct {
	Client     *s3.Client
	BucketName string
}

func (s *S3Storage) UploadFile(id string, data []byte) (string, error) {
	key := fmt.Sprintf("documents/%s", id)

	_, err := s.Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
		Body:   bytes.NewReader(data),
	})

	if err != nil {
		return "", err
	}

	// Return an S3-style URL
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", s.BucketName, key), nil
}

func (s *S3Storage) DownloadFile(id string) ([]byte, error) {
	key := fmt.Sprintf("documents/%s", id)
	out, err := s.Client.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}

	defer out.Body.Close()
	return io.ReadAll(out.Body)
}
