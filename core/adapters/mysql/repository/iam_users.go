package repository

import (
	"fmt"
	"iam-examples-go/core/domain"
	"iam-examples-go/core/ports/repository"

	"gorm.io/gorm"
)

type IAMUsersRepository struct {
	db *gorm.DB
}

func NewIAMUsersRepository(db *gorm.DB) repository.IAMUsersRepository {
	return &IAMUsersRepository{db: db}
}
func (r IAMUsersRepository) All() ([]domain.IAMUsers, error) {
	var users []domain.IAMUsers
	result := r.db.
		Model(&domain.IAMUsers{}).
		Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching IAM Users:", result.Error)
		return nil, result.Error
	}

	return users, nil
}
