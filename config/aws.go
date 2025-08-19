package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type AwsConfig struct {
	Region          string
	S3Bucket        string
	S3Endpoint      string
	AccessKeyID     string
	SecretAccessKey string
}

func NewAwsConfig() *AwsConfig {
	return &AwsConfig{
		Region:          utils.GetEnv("AWS_REGION", ""),
		S3Bucket:        utils.GetEnv("AWS_S3_BUCKET", ""),
		S3Endpoint:      utils.GetEnv("AWS_S3_ENDPOINT", ""),
		AccessKeyID:     utils.GetEnv("AWS_ACCESS_KEY_ID", ""),
		SecretAccessKey: utils.GetEnv("AWS_SECRET_ACCESS_KEY", ""),
	}
}
