package form

import (
	"mime/multipart"
)

type Avatar struct {
	RandomID   string                `form:"randomId" validate:"required"`
	AvatarFile *multipart.FileHeader `form:"avatar" validate:"file_required,file_max_size=2MB,file_types=image/png,image/jpeg"`
}
