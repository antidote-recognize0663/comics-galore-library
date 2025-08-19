package config

import (
	"github.com/joho/godotenv"
	"log"
)

type Config struct {
	NeonDB           *DbConfig
	AWS              *AwsConfig
	Tusd             *TusdConfig
	Redis            *RedisConfig
	ImageDefaults    *ImageConfig
	ImagesAssets     *ImageAssets
	Appwrite         *AppwriteConfig
	Application      *ApplicationConfig
	NowPayments      *NowPaymentsConfig
	CloudflareR2     *CloudflareR2Config
	CloudflareImages *CloudflareImagesConfig
}

func NewConfig(envFiles ...string) *Config {
	if len(envFiles) > 0 {
		err := godotenv.Load(envFiles...)
		if err != nil {
			log.Printf("Warning: Error loading .env file(s): %v", err)
		}
	}
	cfg := &Config{
		AWS:              NewAwsConfig(),
		NeonDB:           NewDbConfig(),
		Tusd:             NewTusdConfig(),
		Redis:            NewRedisConfig(),
		ImagesAssets:     NewImageAssets(),
		CloudflareR2:     NewCloudflareR2(),
		ImageDefaults:    NewImageConfig(),
		Appwrite:         NewAppwriteConfig(),
		NowPayments:      NewNowPaymentsConfig(),
		CloudflareImages: NewCloudflareImages(),
		Application:      NewApplicationConfig(),
	}
	log.Println("Configuration loaded successfully.")
	return cfg
}
