package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/json"
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"sort"
	"strings"
)

func CheckIPNRequestIsValid(receivedHMAC string, receivedData []byte, ipnSecret string) (bool, string, *model.PaymentStatus) {
	errorMsg := "Unknown error"
	authOK := false
	var paymentStatus model.PaymentStatus

	// Parse JSON request data into struct
	err := json.Unmarshal(receivedData, &paymentStatus)
	if err != nil {
		return false, "Invalid JSON format", nil
	}

	// Convert JSON to map for sorting
	var paymentMap map[string]interface{}
	err = json.Unmarshal(receivedData, &paymentMap)
	if err != nil {
		return false, "Error parsing JSON for sorting", nil
	}

	// Sort the request data keys
	sortedKeys := make([]string, 0, len(paymentMap))
	for key := range paymentMap {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	// Create a sorted JSON string
	sortedRequestData := make(map[string]interface{})
	for _, key := range sortedKeys {
		sortedRequestData[key] = paymentMap[key]
	}
	sortedJSON, err := json.Marshal(sortedRequestData)
	if err != nil {
		return false, "Error creating sorted JSON", nil
	}

	// Compute HMAC hash
	h := hmac.New(sha512.New, []byte(ipnSecret))
	h.Write(sortedJSON)
	computedHMAC := fmt.Sprintf("%x", h.Sum(nil))

	// Compare computed HMAC with received HMAC (case-insensitive)
	if strings.EqualFold(computedHMAC, receivedHMAC) {
		authOK = true
	} else {
		errorMsg = "HMAC signature does not match"
	}

	return authOK, errorMsg, &paymentStatus
}
