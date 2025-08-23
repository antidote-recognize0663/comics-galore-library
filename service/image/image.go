package image

import (
	"errors"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
)

type Image interface {
	View(secret string, fileId string, bucketId string) (*[]byte, error)
	Preview(secret string, fileId string, bucketId string, width int, height int, quality int, gravity string) (*[]byte, error)
}

type image struct {
	endpoint  string
	bucketID  string
	projectID string
}

func NewImageService(opts ...Option) Image {
	cfg := &Config{
		endpoint:  "https://fra.cloud.appwrite.io/v1",
		projectID: "6510a59f633f9d57fba2",
		bucketID:  "68574e890011f6c911c3",
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return &image{
		endpoint:  cfg.endpoint,
		bucketID:  cfg.bucketID,
		projectID: cfg.projectID,
	}
}

func NewImageWithConfig(cfg *config.Config) Image {
	return &image{
		endpoint:  cfg.Appwrite.Endpoint,
		bucketID:  cfg.Appwrite.BucketIDImages,
		projectID: cfg.Appwrite.ProjectID,
	}
}

func (i *image) View(secret string, fileId string, bucketId string) (*[]byte, error) {
	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}
	return appwrite.NewStorage(*utils.NewSessionClient(secret, utils.WithProject(i.projectID), utils.WithEndpoint(i.endpoint))).GetFileView(bucketId, fileId)
}

func (i *image) Preview(secret string, fileId string, bucketId string, width int, height int, quality int, gravity string) (*[]byte, error) {
	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}
	if bucketId == "" {
		return nil, errors.New("bucketId can not be empty")
	}
	storage := appwrite.NewStorage(*utils.NewSessionClient(secret, utils.WithProject(i.projectID), utils.WithEndpoint(i.endpoint)))
	return storage.GetFilePreview(bucketId, fileId, storage.WithGetFilePreviewWidth(width),
		storage.WithGetFilePreviewHeight(height),
		storage.WithGetFilePreviewQuality(quality),
		storage.WithGetFilePreviewGravity(gravity))
}

func WithEndpoint(endpoint string) Option {
	return func(config *Config) {
		config.endpoint = endpoint
	}
}

func WithProjectID(projectID string) Option {
	return func(config *Config) {
		config.projectID = projectID
	}
}

func WithBucketID(bucketID string) Option {
	return func(config *Config) {
		config.bucketID = bucketID
	}
}

type Config struct {
	endpoint  string
	bucketID  string
	projectID string
}

type Option func(*Config)
