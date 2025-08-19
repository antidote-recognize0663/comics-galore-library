package form

import (
	"mime/multipart"
)

type UploadRequest struct {
	Title       string                  `form:"title" validate:"required"`
	Author      string                  `form:"author" validate:"required"`
	Category    string                  `form:"category" validate:"required"`
	Description string                  `form:"description"`
	Cover       *multipart.FileHeader   `form:"cover" validate:"file_required,file_types=image/png;image/jpeg;image/jpg;image/webp"`
	Previews    []*multipart.FileHeader `form:"previews[]" validate:"required,gt=0,dive,file_required,file_types=image/png;image/jpeg;image/jpg;image/webp"`
	Archives    []*multipart.FileHeader `form:"archives[]" validate:"required,gt=0,dive,file_types=application/vnd.comicbook+zip;application/vnd.comicbook-rar;application/zip;application/vnd.rar;application/pdf"`
}

func NewUploadRequest(form *multipart.Form) *UploadRequest {
	request := &UploadRequest{}
	if covers, ok := form.File["cover"]; ok && len(covers) > 0 {
		request.Cover = covers[0]
	}
	if previews, ok := form.File["previews[]"]; ok {
		request.Previews = previews
	}
	if archives, ok := form.File["archives[]"]; ok {
		request.Archives = archives
	}
	if titles, ok := form.Value["title"]; ok && len(titles) > 0 {
		request.Title = titles[0]
	}
	if authors, ok := form.Value["author"]; ok && len(authors) > 0 {
		request.Author = authors[0]
	}
	if categories, ok := form.Value["category"]; ok && len(categories) > 0 {
		request.Category = categories[0]
	}
	if descriptions, ok := form.Value["description"]; ok && len(descriptions) > 0 {
		request.Description = descriptions[0]
	}
	return request
}
