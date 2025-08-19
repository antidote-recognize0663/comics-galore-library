package form

import (
	"mime/multipart"
)

type ProfileEmail struct {
	RandomID string `form:"randomId" validate:"required"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=7,password"`
}

type ProfilePrefs struct {
	AvatarId string `form:"AvatarId"`
	Tumblr   string `form:"TumblrBlog"`
	Twitter  string `form:"TwitterHandle"`
	Facebook string `form:"FacebookProfile"`
}

type ProfileAvatar struct {
	AvatarFile *multipart.FileHeader `form:"avatar"`
}

type ProfilePassword struct {
	RandomID    string `form:"randomId" validate:"required"`
	OldPassword string `form:"old_password" validate:"required,min=7,password"`
	NewPassword string `form:"new_password" validate:"required,min=7,password"`
}

type SocialPrefs struct {
	Tumblr   string `form:"TumblrBlog"`
	Twitter  string `form:"TwitterHandle"`
	Facebook string `form:"FacebookProfile"`
}
