package model

import "time"

type CloudflareImageResponse struct {
	Result   CloudflareImageResult `json:"result"`
	Success  bool                  `json:"success"`
	Errors   []CloudflareError     `json:"errors"`
	Messages []CloudflareMessage   `json:"messages"`
}

type CloudflareImageResult struct {
	ID                string            `json:"id"`
	Filename          string            `json:"filename"`
	Meta              map[string]string `json:"meta"`
	RequireSignedURLs bool              `json:"requireSignedURLs"`
	Uploaded          time.Time         `json:"uploaded"`
	Variants          []string          `json:"variants"`
}

type CloudflareError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
type CloudflareMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type CloudflareListImagesResponse struct {
	Errors   []CloudflareError   `json:"errors"`
	Messages []CloudflareMessage `json:"messages"`
	Success  bool                `json:"success"`
	Result   ListImagesResult    `json:"result"`
}

type ListImagesResult struct {
	Images []CloudflareImageResult `json:"images"`
}
