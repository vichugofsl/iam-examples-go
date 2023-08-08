package repository

import (
	"fmt"
	"iam-examples-go/core/domain"
	"iam-examples-go/core/ports/repository"

	"gorm.io/gorm"
)

type IAMUserKeysRepository struct {
	db *gorm.DB
}

func NewIAMUserKeysRepository(db *gorm.DB) repository.IAMUserKeysRepository {
	return &IAMUserKeysRepository{db: db}
}

func (r IAMUserKeysRepository) All() ([]domain.IAMUserKeys, error) {
	var userKeys []domain.IAMUserKeys
	result := r.db.
		Model(&domain.IAMUserKeys{}).
		Find(&userKeys)
	if result.Error != nil {
		fmt.Println("Error fetching IAM User Keys:", result.Error)
		return nil, result.Error
	}

	return userKeys, nil
}
