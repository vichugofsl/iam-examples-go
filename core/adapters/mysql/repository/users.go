package repository

import (
	"fmt"
	"iam-examples-go/core/domain"
	"iam-examples-go/core/ports/repository"
	"log"

	"gorm.io/gorm"
)

type UsersRepository struct {
	db *gorm.DB
}

func NewUsersRepository(db *gorm.DB) repository.UsersRepository {
	return &UsersRepository{db: db}
}

func (r UsersRepository) CreateNewTable() error {
	return r.db.AutoMigrate(&domain.IAMUsers{})
}

func (r UsersRepository) ExtractIAMUsers() error {
	var users []domain.Users
	users, err := r.List()

	if err != nil {
		return err
	}

	log.Printf("It will start to insert the users")

	for _, user := range users {
		iamUser := domain.IAMUsers{
			OldId:     user.ID,
			Name:      user.Name,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Email:     user.Email,
			Password:  user.Password,
			Activated: user.Activated,
			Token:     user.Activated,
		}
		result := r.db.Create(&iamUser)

		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}

func (r UsersRepository) List() ([]domain.Users, error) {
	var users []domain.Users
	result := r.db.
		Model(&domain.Users{}).
		Joins("JOIN user_keys ON user_keys.user_id = users.id").
		Where("deleted_at IS NULL").
		Group("users.id").
		Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching Users:", result.Error)
		return nil, result.Error
	}

	return users, nil
}

// This method returns the total count of records in the Users table.
func (r UsersRepository) Total() (int64, error) {
	var count int64
	result := r.db.
		Model(&domain.Users{}).
		Joins("JOIN user_keys ON user_keys.user_id = users.id").
		Where("deleted_at IS NULL").
		Group("users.id").
		Count(&count)
	if result.Error != nil {
		fmt.Println("Error counting Users:", result.Error)
		return 0, result.Error
	}

	return count, nil
}
