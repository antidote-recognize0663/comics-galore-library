package image

import (
	"errors"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
)

type Avatar interface {
	Avatar(secret string, fileId string, width int, height int, quality int, gravity string) (*[]byte, error)
}

type avatar struct {
	endpoint  string
	bucketID  string
	projectID string
}

func NewAvatar(opts ...Option) Avatar {
	cfg := &Config{
		endpoint:  "https://fra.cloud.appwrite.io/v1",
		projectID: "6512130e80992b6c3e11",
		bucketID:  "651b3476e4b9da11935f",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &avatar{
		endpoint:  cfg.endpoint,
		bucketID:  cfg.bucketID,
		projectID: cfg.projectID,
	}
}

func NewAvatarWithConfig(cfg *config.Config) Avatar {
	return &avatar{
		endpoint:  cfg.Appwrite.Endpoint,
		bucketID:  cfg.Appwrite.BucketIDAvatars,
		projectID: cfg.Appwrite.ProjectID,
	}
}

func (i *avatar) Avatar(secret string, fileId string, width int, height int, quality int, gravity string) (*[]byte, error) {
	if fileId == "" {
		return nil, errors.New("fileId can not be empty")
	}
	storage := appwrite.NewStorage(*utils.NewSessionClient(secret, utils.WithProject(i.projectID), utils.WithEndpoint(i.endpoint)))
	return storage.GetFilePreview(i.bucketID, fileId, storage.WithGetFilePreviewWidth(width),
		storage.WithGetFilePreviewHeight(height),
		storage.WithGetFilePreviewQuality(quality),
		storage.WithGetFilePreviewGravity(gravity))
}
