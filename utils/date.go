package utils

import (
	"fmt"
	"time"
)

func ParseDateTimeToUnix(dateTime, layout string) (int64, error) {
	if layout == "" {
		return 0, fmt.Errorf("layout must not be empty")
	}
	parsedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		return 0, fmt.Errorf("failed to parse datetime: %w", err)
	}
	return parsedTime.Unix(), nil
}

func ParseDateTime(dateTime, layout string) (time.Time, error) {
	if layout == "" {
		return time.Time{}, fmt.Errorf("layout must not be empty")
	}
	parsedTime, err := time.Parse(layout, dateTime)
	if err != nil {
		return time.Time{}, fmt.Errorf("failed to parse datetime: %w", err)
	}
	return parsedTime, nil
}

func FormatDateTime(dateTime string, optionalLayouts ...string) (string, error) {
	var inputLayout, outputLayout string
	if len(optionalLayouts) > 0 && optionalLayouts[0] != "" {
		inputLayout = optionalLayouts[0]
	} else {
		inputLayout = time.RFC3339
	}
	if len(optionalLayouts) > 1 && optionalLayouts[1] != "" {
		outputLayout = optionalLayouts[1]
	} else {
		outputLayout = "02/01/2006 15:04:05"
	}
	parsedTime, err := time.Parse(inputLayout, dateTime)
	if err != nil {
		return "", fmt.Errorf("failed to parse datetime: %w", err)
	}
	return parsedTime.Format(outputLayout), nil
}
