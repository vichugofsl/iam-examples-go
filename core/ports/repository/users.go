package repository

import (
	"iam-examples-go/core/domain"
)

type UsersRepository interface {
	CreateNewTable() error
	ExtractIAMUsers() error
	List() ([]domain.Users, error)
	Total() (int64, error)
}

type IAMUsersRepository interface {
	All() ([]domain.IAMUsers, error)
}
