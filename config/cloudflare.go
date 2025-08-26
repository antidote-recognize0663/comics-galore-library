package config

type CloudflareImagesConfig struct {
	ApiKey    string
	AccountID string
	ImagesURL string
}

func NewCloudflareImages() *CloudflareImagesConfig {
	return &CloudflareImagesConfig{
		ApiKey:    GetEnv("CLOUDFLARE_API_KEY", ""),
		AccountID: GetEnv("CLOUDFLARE_ACCOUNT_ID", ""),
		ImagesURL: GetEnv("CLOUDFLARE_IMAGES_URL", ""),
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
		Bucket:          GetEnv("CLOUDFLARE_R2_BUCKET", ""),
		Endpoint:        GetEnv("CLOUDFLARE_R2_ENDPOINT", ""),
		AccessKey:       GetEnv("CLOUDFLARE_R2_ACCESS_KEY", ""),
		SecretAccessKey: GetEnv("CLOUDFLARE_R2_SECRET_ACCESS_KEY", ""),
	}
}
