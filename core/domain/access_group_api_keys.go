package domain

import (
	"time"

	"gorm.io/gorm"
)

type AccessGroupApiKeys struct {
	AccessGroupId uint
	KeyId         uint
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

func (AccessGroupApiKeys) TableName() string {
	return "access_group_api_keys"
}

type IAMAccessGroupApiKeys struct {
	gorm.Model
	AccessGroupId uint
	UserKeyId     uint
	UserKey       IAMUserKeys
}

func (IAMAccessGroupApiKeys) TableName() string {
	return "iam_access_group_api_keys"
}
