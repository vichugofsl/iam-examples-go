package repository

import (
	"fmt"
	"iam-examples-go/core/domain"
	"iam-examples-go/core/ports/repository"
	"log"

	"gorm.io/gorm"
)

type UserKeysRepository struct {
	db *gorm.DB
}

func NewUserKeysRepository(db *gorm.DB) repository.UserKeysRepository {
	return &UserKeysRepository{db: db}
}

func (r UserKeysRepository) CreateNewTable() error {
	return r.db.AutoMigrate(&domain.IAMUserKeys{})
}

func (r UserKeysRepository) ExtractIAMUserKeys() error {
	iamUserRepository := NewIAMUsersRepository(r.db)
	users, _ := iamUserRepository.All()
	userMap := make(map[uint]uint)

	for _, user := range users {
		userMap[user.OldId] = user.ID
	}

	var userKeys []domain.UserKeys
	userKeys, err := r.List()

	if err != nil {
		return err
	}

	log.Printf("It will start to insert the user Keys")

	for _, userKey := range userKeys {
		newUserId, hasOldId := userMap[userKey.UserId]
		if hasOldId {
			iamUser := domain.IAMUserKeys{
				OldId:       userKey.ID,
				UserId:      newUserId,
				Key:         userKey.Key,
				Name:        userKey.Name,
				Description: userKey.Description,
			}
			result := r.db.Create(&iamUser)

			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

func (r UserKeysRepository) List() ([]domain.UserKeys, error) {
	var users []domain.UserKeys
	result := r.db.
		Model(&domain.UserKeys{}).
		Joins("JOIN users ON user_keys.user_id = users.id").
		Find(&users)
	if result.Error != nil {
		fmt.Println("Error fetching Users:", result.Error)
		return nil, result.Error
	}

	return users, nil
}

// This method returns the total count of records in the Users table.
func (r UserKeysRepository) Total() (int64, error) {
	var count int64
	result := r.db.
		Model(&domain.UserKeys{}).
		Joins("JOIN users ON user_keys.user_id = users.id").
		Count(&count)
	if result.Error != nil {
		fmt.Println("Error counting Users:", result.Error)
		return 0, result.Error
	}

	return count, nil
}
