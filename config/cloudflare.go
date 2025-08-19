package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type CloudflareImagesConfig struct {
	ApiKey    string
	AccountID string
	ImagesURL string
}

func NewCloudflareImages() *CloudflareImagesConfig {
	return &CloudflareImagesConfig{
		ApiKey:    utils.GetEnv("CLOUDFLARE_API_KEY", ""),
		AccountID: utils.GetEnv("CLOUDFLARE_ACCOUNT_ID", ""),
		ImagesURL: utils.GetEnv("CLOUDFLARE_IMAGES_URL", ""),
	}
}

type CloudflareR2Config struct {
	Bucket          string
	Endpoint        string
	AccessKey       string
	SecretAccessKey string
}

func NewCloudflareR2() *CloudflareR2Config {
	return &CloudflareR2Config{
		Bucket:          utils.GetEnv("CLOUDFLARE_R2_BUCKET", ""),
		Endpoint:        utils.GetEnv("CLOUDFLARE_R2_ENDPOINT", ""),
		AccessKey:       utils.GetEnv("CLOUDFLARE_R2_ACCESS_KEY", ""),
		SecretAccessKey: utils.GetEnv("CLOUDFLARE_R2_SECRET_ACCESS_KEY", ""),
	}
}
