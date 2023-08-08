package domain

import (
	"time"

	"gorm.io/gorm"
)

type UserKeys struct {
	ID          uint `gorm:"primarykey"`
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

type IAMUserKeys struct {
	gorm.Model
	OldId       uint
	UserId      uint
	Key         string
	Name        string
	Description string
	User        IAMUsers
}

func (IAMUserKeys) TableName() string {
	return "iam_user_keys"
}
