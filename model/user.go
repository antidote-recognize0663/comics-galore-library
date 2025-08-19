package model

import "github.com/appwrite/sdk-for-go/models"

type Preferences struct {
	AvatarId string `json:"avatar_id,omitempty"`
	Twitter  string `json:"twitter,omitempty"`
	Facebook string `json:"facebook,omitempty"`
	Tumblr   string `json:"tumblr,omitempty"`
}

type User struct {
	*models.User
	*Preferences
}
