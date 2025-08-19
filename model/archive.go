package model

import (
	"github.com/appwrite/sdk-for-go/models"
)

type Archive struct {
	*models.Document
	*ArchiveData
}

type ArchiveData struct {
	Name      string `json:"name"`
	Key       string `json:"key"`
	Size      int64  `json:"size"`
	Pages     int    `json:"pages"`
	FileID    string `json:"file_id"`
	Bucket    string `json:"bucket"`
	MimeType  string `json:"mime_type"`
	Downloads int64  `json:"downloads"`
}

type ArchiveList struct {
	*models.DocumentList
	Archives []Archive `json:"archives"`
}
