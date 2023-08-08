package repository

import (
	"iam-examples-go/core/domain"
)

type UserKeysRepository interface {
	CreateNewTable() error
	ExtractIAMUserKeys() error
	List() ([]domain.UserKeys, error)
	Total() (int64, error)
}

type IAMUserKeysRepository interface {
	All() ([]domain.IAMUserKeys, error)
}
