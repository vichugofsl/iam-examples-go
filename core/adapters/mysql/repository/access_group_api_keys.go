package repository

import (
	"fmt"
	"iam-examples-go/core/domain"
	"iam-examples-go/core/ports/repository"
	"log"

	"gorm.io/gorm"
)

type AccessGroupApiKeysRepository struct {
	db *gorm.DB
}

func NewAccessGroupApiKeysRepository(db *gorm.DB) repository.AccessGroupApiKeysRepository {
	return &AccessGroupApiKeysRepository{db: db}
}

func (r AccessGroupApiKeysRepository) CreateNewTable() error {
	return r.db.AutoMigrate(&domain.IAMAccessGroupApiKeys{})
}

func (r AccessGroupApiKeysRepository) ExtractIAMAccessGroupApiKeys() error {
	iamUserKeysRepository := NewIAMUserKeysRepository(r.db)
	userKeys, _ := iamUserKeysRepository.All()
	userMap := make(map[uint]uint)

	for _, userKey := range userKeys {
		userMap[userKey.OldId] = userKey.ID
	}

	var accessGroupList []domain.AccessGroupApiKeys
	accessGroupList, err := r.List()

	if err != nil {
		return err
	}

	log.Printf("It will start to insert the access group api Keys")

	for _, accessGroup := range accessGroupList {
		newUserKeyId, hasAccessGroup := userMap[accessGroup.KeyId]
		if hasAccessGroup {
			iamAccessGroup := domain.IAMAccessGroupApiKeys{
				UserKeyId:     newUserKeyId,
				AccessGroupId: accessGroup.AccessGroupId,
			}
			result := r.db.Create(&iamAccessGroup)

			if result.Error != nil {
				return result.Error
			}
		}
	}
	return nil
}

func (r AccessGroupApiKeysRepository) List() ([]domain.AccessGroupApiKeys, error) {
	var groups []domain.AccessGroupApiKeys
	result := r.db.
		Model(&domain.AccessGroupApiKeys{}).
		Joins("JOIN user_keys ON user_keys.id = access_group_api_keys.key_id").
		Joins("JOIN users ON user_keys.user_id = users.id").
		Where("users.deleted_at IS NULL").
		Group("access_group_api_keys.key_id, access_group_api_keys.access_group_id").
		Find(&groups)
	if result.Error != nil {
		fmt.Println("Error fetching Users:", result.Error)
		return nil, result.Error
	}

	return groups, nil
}

// This method returns the total count of records in the Users table.
func (r AccessGroupApiKeysRepository) Total() (int64, error) {
	var count int64
	result := r.db.
		Model(&domain.AccessGroupApiKeys{}).
		Joins("JOIN user_keys ON user_keys.id = access_group_api_keys.key_id").
		Joins("JOIN users ON user_keys.user_id = users.id").
		Where("users.deleted_at IS NULL").
		Group("access_group_api_keys.key_id, access_group_api_keys.access_group_id").
		Count(&count)
	if result.Error != nil {
		fmt.Println("Error counting Users:", result.Error)
		return 0, result.Error
	}

	return count, nil
}
