package utils

import (
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"time"
)

func GetExpirationDate(startTime time.Time, amount float64, durationMap map[float64]model.Duration) (string, error) {
	duration, exists := durationMap[amount]
	if !exists {
		return "", fmt.Errorf("invalid amount: no duration configured for %.2f", amount)
	}
	expirationDate := startTime.AddDate(duration.Years, duration.Months, duration.Days)
	return expirationDate.Format(time.RFC3339), nil
}

func BuildQRCodeURL(currency, address string, amount float64, id string) string {
	const defaultSize = "200"
	const defaultMargin = "0"
	return fmt.Sprintf("/qrcode/%s/%s?amount=%.2f&label=%s&size=%s&margin=%s",
		currency, address, amount, id, defaultSize, defaultMargin)
}
