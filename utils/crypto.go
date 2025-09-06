package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
)

// Encrypt - Encrypts a string using AES and returns a Base64-encoded encrypted string
func Encrypt(plainText, key string) (string, error) {
	// AES requires the key to have a length of 16, 24, or 32 bytes
	// Pad the key if necessary to make it exactly 32 bytes
	for len(key) < 32 {
		key += "0"
	}
	key = key[:32]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Convert the plaintext string into a byte slice
	plainTextBytes := []byte(plainText)

	// Generate a random initialization vector (IV)
	ciphertext := make([]byte, aes.BlockSize+len(plainTextBytes))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Use CBC mode and encrypt the plaintext
	mode := cipher.NewCBCEncrypter(block, iv)
	paddedPlaintext := padPKCS7(plainTextBytes, aes.BlockSize)
	mode.CryptBlocks(ciphertext[aes.BlockSize:], paddedPlaintext)

	// Return the Base64-encoded ciphertext
	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// Decrypt - Decrypts a Base64-encoded encrypted string into the original string
func Decrypt(encryptedText, key string) (string, error) {
	// Pad the key if necessary to make it exactly 32 bytes
	for len(key) < 32 {
		key += "0"
	}
	key = key[:32]

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	// Decode the base64-encoded string
	ciphertext, err := base64.URLEncoding.DecodeString(encryptedText)
	if err != nil {
		return "", err
	}

	// Get the initialization vector (IV) from the beginning of the ciphertext
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Use CBC mode and decrypt the ciphertext
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)

	// Remove padding and return the original plaintext as a string
	plainText := unpadPKCS7(ciphertext)
	return string(plainText), nil
}

// padPKCS7 - Pads data using the PKCS#7 standard
func padPKCS7(data []byte, blockSize int) []byte {
	padding := blockSize - len(data)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padText...)
}

// unpadPKCS7 - Removes PKCS#7 padding
func unpadPKCS7(data []byte) []byte {
	length := len(data)
	padding := int(data[length-1])
	return data[:length-padding]
}
