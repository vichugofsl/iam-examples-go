package domain

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	V2Id      uint
	Name      string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Activated string
	Token     string
}

func (Users) TableName() string {
	return "users"
}

type IAMUsers struct {
	gorm.Model
	OldId     uint
	Name      string
	FirstName string
	LastName  string
	Email     string
	Password  string
	Activated string
	Token     string
}

func (IAMUsers) TableName() string {
	return "iam_users"
}
