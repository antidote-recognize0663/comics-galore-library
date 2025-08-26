package config

type AwsConfig struct {
	Region          string
	S3Bucket        string
	S3Endpoint      string
	AccessKeyID     string
	SecretAccessKey string
}

func NewAwsConfig() *AwsConfig {
	return &AwsConfig{
		Region:          GetEnv("AWS_REGION", ""),
		S3Bucket:        GetEnv("AWS_S3_BUCKET", ""),
		S3Endpoint:      GetEnv("AWS_S3_ENDPOINT", ""),
		AccessKeyID:     GetEnv("AWS_ACCESS_KEY_ID", ""),
		SecretAccessKey: GetEnv("AWS_SECRET_ACCESS_KEY", ""),
	}
}
