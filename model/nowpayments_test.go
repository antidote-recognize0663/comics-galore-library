package model

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/hex"
	"testing"
)

func TestGenerateCorrectHMAC(t *testing.T) {
	const ipnSecret = "my-super-secret-ipn-key"
	const validPayload = `{"created_at":"2025-08-16T00:30:00Z","order_id":"my-order-987","pay_amount":0.00025,"pay_currency":"btc","payment_id":12345678,"payment_status":"finished","price_amount":10.50,"price_currency":"usd"}`

	canonicalPayload, err := createCanonicalJSON([]byte(validPayload))
	if err != nil {
		t.Fatalf("Failed to create canonical JSON: %v", err)
	}

	// Note: The original payload had the keys in a different order. Our function will sort them.
	// The canonical payload string will be:
	// {"created_at":"2025-08-16T00:30:00Z","order_id":"my-order-987","pay_amount":0.00025,"pay_currency":"btc","payment_id":12345678,"payment_status":"finished","price_amount":10.50,"price_currency":"usd"}

	h := hmac.New(sha512.New, []byte(ipnSecret))
	h.Write([]byte(canonicalPayload))
	correctHMACForTest := hex.EncodeToString(h.Sum(nil))

	t.Logf("The correct HMAC for your test is: %s", correctHMACForTest)
}

func TestValidatePayload(t *testing.T) {
	const ipnSecret = "my-super-secret-ipn-key"
	const validPayload = `{"payment_id": 12345678, "payment_status": "finished", "price_amount": 10.50, "price_currency": "usd", "pay_amount": 0.00025, "pay_currency": "btc", "order_id": "my-order-987", "created_at": "2025-08-16T00:30:00Z"}`
	const correctHMAC = "282306035f763e7f75fb594c8fd484928599988a95cee7e682802b3fd75c5c03bf2bd28e48ba225067338e8622247b4d7129f5cd84fe4424c29f081fac551e37"
	const tamperedPayload = `{"payment_id": 12345678, "payment_status": "finished", "price_amount": 10.51, "price_currency": "usd", "pay_amount": 0.00025, "pay_currency": "btc", "order_id": "my-order-987", "created_at": "2025-08-16T00:30:00Z"}`
	const malformedPayload = `{"payment_id": 12345,`

	testCases := []struct {
		name         string
		hmac         string
		payload      []byte
		secret       string
		expectAuthOK bool
		expectErrMsg string
	}{
		{
			name:         "Valid Signature",
			hmac:         correctHMAC,
			payload:      []byte(validPayload),
			secret:       ipnSecret,
			expectAuthOK: true,
		},
		{
			name:         "Invalid Signature (Wrong HMAC)",
			hmac:         "a_completely_wrong_hmac_string",
			payload:      []byte(validPayload),
			secret:       ipnSecret,
			expectAuthOK: false,
			expectErrMsg: "HMAC signature does not match",
		},
		{
			name:         "Invalid Signature (Tampered Payload)",
			hmac:         correctHMAC,
			payload:      []byte(tamperedPayload),
			secret:       ipnSecret,
			expectAuthOK: false,
			expectErrMsg: "HMAC signature does not match",
		},
		{
			name:         "Malformed JSON Payload",
			hmac:         correctHMAC,
			payload:      []byte(malformedPayload),
			secret:       ipnSecret,
			expectAuthOK: false,
			expectErrMsg: "Invalid JSON format",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			authOK, errorMsg, _ := ValidatePayload(tc.hmac, tc.payload, tc.secret)

			if authOK != tc.expectAuthOK {
				t.Errorf("Expected authOK to be %v, but got %v", tc.expectAuthOK, authOK)
			}

			if !tc.expectAuthOK && errorMsg != tc.expectErrMsg {
				t.Errorf("Expected error message '%s', but got '%s'", tc.expectErrMsg, errorMsg)
			}
		})
	}
}
