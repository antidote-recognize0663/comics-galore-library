package backblaze

import (
	"bytes"
	"context"
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"io"
	"log"
)

type FileStorage interface {
	GetFile(key string) ([]byte, error)
	PutFile(key string, data []byte) error
	DeleteFile(key string) error
	Reset() error
}

type fileStorage struct {
	client *minio.Client
	bucket string
}

func (a fileStorage) GetFile(key string) ([]byte, error) {
	reader, err := a.client.GetObject(context.Background(), a.bucket, key, minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}
	defer func(reader *minio.Object) {
		if err := reader.Close(); err != nil {
			log.Printf("Failed to close reader: %v", err)
		}
	}(reader)
	var buffer bytes.Buffer
	if _, err := io.Copy(&buffer, reader); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (a fileStorage) PutFile(key string, data []byte) error {
	contentReader := bytes.NewReader(data)
	info, err := a.client.PutObject(
		context.Background(),
		a.bucket,
		key,
		contentReader,
		contentReader.Size(),
		minio.PutObjectOptions{
			ContentType: "application/octet-stream",
		},
	)
	if err != nil {
		return fmt.Errorf("failed to upload data to S3: %w", err)
	}
	log.Printf("Uploaded object %s to S3 as %s, size: %d bytes", key, info.Key, info.Size)
	return nil
}

func (a fileStorage) DeleteFile(key string) error {
	err := a.client.RemoveObject(context.Background(), a.bucket, key, minio.RemoveObjectOptions{})
	if err != nil {
		return fmt.Errorf("failed to delete object %s from bucket %s: %w", key, a.bucket, err)
	}
	log.Printf("Successfully deleted object %s from bucket %s", key, a.bucket)
	return nil
}

func (a fileStorage) Reset() error {
	objectCh := a.client.ListObjects(context.Background(), a.bucket, minio.ListObjectsOptions{
		Recursive: true,
	})
	for object := range objectCh {
		if object.Err != nil {
			return fmt.Errorf("error listing objects in bucket %s: %w", a.bucket, object.Err)
		}
		err := a.client.RemoveObject(context.Background(), a.bucket, object.Key, minio.RemoveObjectOptions{})
		if err != nil {
			return fmt.Errorf("failed to delete object %s: %w", object.Key, err)
		}
		log.Printf("Deleted object %s from bucket %s", object.Key, a.bucket)
	}
	log.Printf("Successfully reset the bucket %s", a.bucket)
	return nil
}

func NewArchiveWithConfig(config config.Config) FileStorage {
	minioClient, err := minio.New(config.AWS.S3Endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(config.AWS.AccessKeyID, config.AWS.SecretAccessKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &fileStorage{
		bucket: config.AWS.S3Bucket,
		client: minioClient,
	}
}

type Config struct {
	bucket    string
	endpoint  string
	secretKey string
	accessKey string
}

type Option func(*Config)

func NewArchive(opts ...Option) FileStorage {
	cfg := &Config{
		bucket:   "comics-galore",
		endpoint: "s3.us-east-005.backblazeb2.com",
	}
	for _, option := range opts {
		option(cfg)
	}
	minioClient, err := minio.New(cfg.endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(cfg.accessKey, cfg.secretKey, ""),
		Secure: true,
	})
	if err != nil {
		log.Fatalln(err)
	}
	return &fileStorage{
		bucket: cfg.bucket,
		client: minioClient,
	}
}

func WithBucket(bucket string) Option {
	return func(cfg *Config) {
		cfg.bucket = bucket
	}
}

func WithEndpoint(endpoint string) Option {
	return func(cfg *Config) {
		cfg.endpoint = endpoint
	}
}

func WithSecretKey(secretKey string) Option {
	return func(cfg *Config) {
		cfg.secretKey = secretKey
	}
}

func WithAccessKey(accessKey string) Option {
	return func(cfg *Config) {
		cfg.accessKey = accessKey
	}
}
