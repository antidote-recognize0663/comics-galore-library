package config

import (
	"github.com/antidote-recognize0663/comics-galore-library/config/utils"
	"log"
)

type RedisConfig struct {
	Endpoint string
	Password string
	DB       int
}

func NewRedisConfig() *RedisConfig {
	var parseErr error
	config := &RedisConfig{
		Endpoint: utils.GetEnv("REDIS_ENDPOINT", ""),
		Password: utils.GetEnv("REDIS_PASSWORD", ""),
		DB:       utils.GetIntEnv("REDIS_DB", 0, &parseErr),
	}
	if parseErr != nil {
		log.Printf("error parsing integer environment variables: %w", parseErr)
	}
	return config
}
