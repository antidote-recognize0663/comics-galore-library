package model

import "github.com/appwrite/sdk-for-go/models"

type ImageData struct {
	FileID            string            `json:"file_id"`
	Size              int64             `json:"size"`
	Filename          string            `json:"filename"`
	Variants          []string          `json:"variants"`
	Uploaded          string            `json:"uploaded"`
	Meta              map[string]string `json:"-"`
	RequireSignedURLs bool              `json:"require_signed_urls"`
}

type Image struct {
	*models.Document
	*ImageData
}

type ImageList struct {
	*models.DocumentList
	Images []*Image `json:"documents"`
}
