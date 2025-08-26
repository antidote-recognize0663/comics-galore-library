package model

import (
	"github.com/appwrite/sdk-for-go/models"
)

type PostData struct {
	IsFavorite  bool      `json:"-"`
	Title       string    `json:"title"`
	Author      string    `json:"author"`
	Category    string    `json:"category"`
	UploaderID  string    `json:"uploader_id"`
	Description string    `json:"description"`
	Cover       Image     `json:"cover"`
	Previews    []Image   `json:"previews"`
	Archives    []Archive `json:"archives"`
	Metrics     *Metrics  `json:"metrics"`
}

type Post struct {
	*models.Document
	*PostData
}

type PostList struct {
	*models.DocumentList
	Posts []Post `json:"documents"`
}
