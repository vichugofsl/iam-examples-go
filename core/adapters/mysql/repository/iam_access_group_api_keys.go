package repository

import (
	"fmt"
	"iam-examples-go/core/domain"
	"iam-examples-go/core/ports/repository"

	"gorm.io/gorm"
)

type IAMAccessGroupApiKeysRepository struct {
	db *gorm.DB
}

func NewIAMAccessGroupApiKeysRepository(db *gorm.DB) repository.IAMAccessGroupApiKeysRepository {
	return &IAMAccessGroupApiKeysRepository{db: db}
}

func (r IAMAccessGroupApiKeysRepository) All() ([]domain.IAMAccessGroupApiKeys, error) {
	var users []domain.IAMAccessGroupApiKeys
	result := r.db.
		Model(&domain.IAMAccessGroupApiKeys{}).
		Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching IAM access group API:", result.Error)
		return nil, result.Error
	}

	return users, nil
}

func (r IAMAccessGroupApiKeysRepository) CreateNewTable() error {
	return r.db.AutoMigrate(&domain.IAMAccessGroupApiKeys{})
}
