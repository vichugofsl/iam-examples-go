package repository

import (
	"user-domain-go/core/domain"
)

type UsersRepository interface {
	// CreateTempTable() error
	CreateNewTable() error
	ExtractIAMUsers() error
	List() ([]domain.Users, error)
	Total() (int64, error)
}

type IAMUsersRepository interface {
	InsertBatch(listToInsert []domain.IAMUsers) error
	Update(user *domain.IAMUsers) error
	GetByBibleFileId(bibleFileId int64) ([]domain.IAMUsers, error)
}
