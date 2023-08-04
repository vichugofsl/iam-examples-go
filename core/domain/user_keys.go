package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserKeys struct {
	gorm.Model
	UserId      uint
	Key         string
	Name        string
	Description string
	User        Users
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (UserKeys) TableName() string {
	return "user_keys"
}
