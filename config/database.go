package config

type DbConfig struct {
	Host     string
	User     string
	Port     string
	DbName   string
	Password string
}

func NewDbConfig() *DbConfig {
	return &DbConfig{
		Host:     GetEnv("NEON_DB_HOST", ""),
		User:     GetEnv("NEON_DB_USER", ""),
		Port:     GetEnv("NEON_DB_PORT", "5432"),
		DbName:   GetEnv("NEON_DB_DBNAME", ""),
		Password: GetEnv("NEON_DB_PASSWORD", ""),
	}
}
