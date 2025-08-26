package config

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func GetEnv(key, fallback string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return fallback
}

func GetIntEnv(key string, fallback int, errAcc *error) int {
	strValue := GetEnv(key, "")
	if strValue == "" {
		return fallback
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		errMessage := fmt.Errorf("invalid integer value for env var %s ('%s'): %w", key, strValue, err)
		if errAcc != nil {
			if *errAcc == nil {
				*errAcc = errMessage
			} else {
				log.Printf("Warning: %v (previous error: %v)", errMessage, *errAcc)
			}
		} else {
			log.Printf("Warning: %v", errMessage)
		}
		return fallback
	}
	return value
}
