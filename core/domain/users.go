package domain

import (
	"time"

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
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}

func (Users) TableName() string {
	return "users"
}

// BibleFileWithGaps is a temp entity which holds information about
// Bible files with gaps in their timestamp sequences.
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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (IAMUsers) TableName() string {
	return "iam_users"
}
