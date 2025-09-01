package image

import (
	"errors"
	"github.com/antidote-recognize0663/comics-galore-library/config"
	"github.com/antidote-recognize0663/comics-galore-library/utils"
	"github.com/appwrite/sdk-for-go/appwrite"
)

type Qrcode interface {
	QR(secret, text string, size int, margin int) (*[]byte, error)
}

type qrcode struct {
	endpoint  string
	projectID string
}

func NewQrcode(opts ...Option) Qrcode {
	cfg := &Config{
		endpoint:  "https://fra.cloud.appwrite.io/v1",
		projectID: "6510a59f633f9d57fba2",
	}
	for _, opt := range opts {
		opt(cfg)
	}
	return &qrcode{
		endpoint:  cfg.endpoint,
		projectID: cfg.projectID,
	}
}

func NewQrcodeWithConfig(cfg *config.Config) Qrcode {
	return &qrcode{
		endpoint:  cfg.Appwrite.Endpoint,
		projectID: cfg.Appwrite.ProjectID,
	}
}

func (i *qrcode) QR(secret, text string, size int, margin int) (*[]byte, error) {
	if text == "" {
		return nil, errors.New("text can not be empty")
	}
	avatars := appwrite.NewAvatars(*utils.NewSessionClient(secret, utils.WithProject(i.projectID), utils.WithEndpoint(i.endpoint)))
	return avatars.GetQR(text, avatars.WithGetQRSize(size), avatars.WithGetQRMargin(margin))
}
