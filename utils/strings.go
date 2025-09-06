package utils

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/url"
	"regexp"
	"strings"
	"time"
)

var seededRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func GenerateRandomString(n int) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // Character pool
	randomString := make([]byte, n)
	for i := range randomString {
		randomIndex := seededRand.Intn(len(letters))
		randomString[i] = letters[randomIndex]
	}
	return string(randomString)
}

func CleanFileName(name string) string {
	re := regexp.MustCompile(`[<>:"/\\|?*\x00-\x1F]`)
	cleanName := re.ReplaceAllString(name, "-")
	cleanName = strings.TrimSpace(strings.ToLower(cleanName))
	return cleanName
}

func GenerateSecureRandomID(length int) (string, error) {
	if length <= 0 {
		return "", fmt.Errorf("length must be greater than 0")
	}
	bytes := make([]byte, length)
	if _, err := rand.Read(bytes); err != nil {
		return "", fmt.Errorf("failed to generate secure random ID: %v", err)
	}
	return hex.EncodeToString(bytes), nil
}

func UpdateQueryParam(originalURL, paramName, newValue string) (string, error) {
	u, err := url.Parse(originalURL)
	if err != nil {
		return "", err
	}
	queryValues := u.Query()
	queryValues.Set(paramName, newValue)
	u.RawQuery = queryValues.Encode()
	return u.String(), nil
}
