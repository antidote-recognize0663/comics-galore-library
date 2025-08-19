package config

import "github.com/antidote-recognize0663/comics-galore-library/config/utils"

type DbConfig struct {
	Host     string
	User     string
	Port     string
	DbName   string
	Password string
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Host:     utils.GetEnv("NEON_DB_HOST", ""),
		User:     utils.GetEnv("NEON_DB_USER", ""),
		Port:     utils.GetEnv("NEON_DB_PORT", "5432"),
		DbName:   utils.GetEnv("NEON_DB_DBNAME", ""),
		Password: utils.GetEnv("NEON_DB_PASSWORD", ""),
	}
}
