package utils

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/antidote-recognize0663/comics-galore-library/model"
	"sort"
	"strings"
)

// createCanonicalJSON takes a raw JSON payload and returns a deterministically sorted,
// compact JSON string, preserving original value formatting.
func createCanonicalJSON(payload []byte) (string, error) {
	// Unmarshal into a map of json.RawMessage to preserve exact value formatting.
	var dataMap map[string]json.RawMessage
	if err := json.Unmarshal(payload, &dataMap); err != nil {
		return "", fmt.Errorf("error parsing JSON for sorting: %w", err)
	}

	// Sort the keys alphabetically.
	sortedKeys := make([]string, 0, len(dataMap))
	for key := range dataMap {
		sortedKeys = append(sortedKeys, key)
	}
	sort.Strings(sortedKeys)

	// Manually build the canonical JSON string from sorted keys and raw values.
	var builder strings.Builder
	builder.WriteString("{")
	for i, key := range sortedKeys {
		if i > 0 {
			builder.WriteString(",")
		}
		// Marshal just the key to get it properly quoted and escaped.
		keyBytes, _ := json.Marshal(key)
		builder.Write(keyBytes)
		builder.WriteString(":")
		// Write the raw value directly, as it's already a valid JSON snippet.
		builder.Write(dataMap[key])
	}
	builder.WriteString("}")

	return builder.String(), nil
}

func IsIpnRequestValid(receivedHMAC string, receivedPayload []byte, ipnSecret string) (bool, string, *model.NowPaymentsIPN) {
	var nowPaymentsIpn model.NowPaymentsIPN
	if err := json.Unmarshal(receivedPayload, &nowPaymentsIpn); err != nil {
		return false, "Invalid JSON format", nil
	}

	// 1. Create the canonical payload using the new helper function.
	canonicalPayload, err := createCanonicalJSON(receivedPayload)
	if err != nil {
		return false, err.Error(), nil
	}

	// 2. Compute the HMAC using the canonical payload.
	h := hmac.New(sha512.New, []byte(ipnSecret))
	h.Write([]byte(canonicalPayload))
	computedHMAC := hex.EncodeToString(h.Sum(nil))

	// 3. Securely compare the signatures.
	if hmac.Equal([]byte(computedHMAC), []byte(receivedHMAC)) {
		return true, "", &nowPaymentsIpn
	}

	return false, "HMAC signature does not match", &nowPaymentsIpn
}
