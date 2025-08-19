package model

import (
	"github.com/google/uuid"
	"time"
)

type Comment struct {
	ID        uuid.UUID  `gorm:"type:uuid;primaryKey"`
	ParentID  *uuid.UUID `gorm:"type:uuid"`
	PostID    string     `gorm:"type:varchar(255);not null"`
	UserID    string     `gorm:"type:varchar(255);not null"`
	Username  string     `gorm:"type:varchar(255);not null"`
	Text      string     `gorm:"type:text;not null"`
	AvatarID  *string    `gorm:"type:varchar(255)"`
	CreatedAt time.Time  `gorm:"autoCreateTime"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime"`
	Replies   []Comment  `gorm:"foreignKey:ParentID;constraint:OnDelete:CASCADE"`
}
