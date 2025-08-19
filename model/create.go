package model

import "github.com/antidote-recognize0663/comics-galore-library/form"

type CreatePost struct {
	Title       string         `json:"title"`
	Author      string         `json:"author"`
	Category    string         `json:"category"`
	Description string         `json:"description"`
	UploaderID  string         `json:"uploader_id"`
	Cover       *ImageData     `json:"cover"`
	Previews    []*ImageData   `json:"previews"`
	Archives    []*ArchiveData `json:"archives"`
}

func NewCreatePost(upload *form.UploadRequest, userId string) *CreatePost {
	return &CreatePost{
		UploaderID:  userId,
		Title:       upload.Title,
		Author:      upload.Author,
		Category:    upload.Category,
		Description: upload.Description,
		Cover:       &ImageData{},
		Previews:    []*ImageData{},
		Archives:    []*ArchiveData{},
	}
}
