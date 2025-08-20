package image

import (
	"errors"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
)

type Image interface {
	View(secret string, fileId string, bucketId string) (*[]byte, error)
	Preview(secret string, fileId string, bucketId string, width int, height int, quality int, gravity string) (*[]byte, error)
	QR(secret, text string, size int, margin int) (*[]byte, error)
	Avatar(secret string, fileId string, width int, height int, quality int, gravity string) (*[]byte, error)
}

type image struct {
	endpoint       string
	bucketID       string
	projectID      string
	avatarBucketID string
}

func NewImageService(opts ...Option) Image {
	config := &Config{
		endpoint:       "https://fra.cloud.appwrite.io/v1",
		projectID:      "6510a59f633f9d57fba2",
		bucketID:       "68574e890011f6c911c3",
		avatarBucketID: "651b3476e4b9da11935f",
	}
	for _, opt := range opts {
		opt(config)
	}

	return &image{
		endpoint:       config.endpoint,
		bucketID:       config.bucketID,
		projectID:      config.projectID,
		avatarBucketID: config.avatarBucketID,
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

func (i *image) QR(secret, text string, size int, margin int) (*[]byte, error) {
	if text == "" {
		return nil, errors.New("text can not be empty")
	}
	avatars := appwrite.NewAvatars(*utils.NewSessionClient(secret, utils.WithProject(i.projectID), utils.WithEndpoint(i.endpoint)))
	return avatars.GetQR(text, avatars.WithGetQRSize(size), avatars.WithGetQRMargin(margin))
}

func (i *image) Avatar(secret string, fileId string, width int, height int, quality int, gravity string) (*[]byte, error) {
	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}
	storage := appwrite.NewStorage(*utils.NewSessionClient(secret, utils.WithProject(i.projectID), utils.WithEndpoint(i.endpoint)))
	return storage.GetFilePreview(i.avatarBucketID, fileId, storage.WithGetFilePreviewWidth(width),
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

func WithAvatarBucketID(bucketID string) Option {
	return func(config *Config) {
		config.avatarBucketID = bucketID
	}
}

type Config struct {
	endpoint       string
	bucketID       string
	projectID      string
	avatarBucketID string
}

type Option func(*Config)
