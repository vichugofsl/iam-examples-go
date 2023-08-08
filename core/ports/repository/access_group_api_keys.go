package repository

import (
	"iam-examples-go/core/domain"
)

type AccessGroupApiKeysRepository interface {
	CreateNewTable() error
	ExtractIAMAccessGroupApiKeys() error
	List() ([]domain.AccessGroupApiKeys, error)
	Total() (int64, error)
}

type IAMAccessGroupApiKeysRepository interface {
	CreateNewTable() error
	All() ([]domain.IAMAccessGroupApiKeys, error)
}
