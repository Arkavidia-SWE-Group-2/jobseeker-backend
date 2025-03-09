package repository

import (
	"jobseeker/internal/entity"

	"gorm.io/gorm"
)

type (
	UserRepository interface {
		Create(tx *gorm.DB, user *entity.User) error
		ExistsByEmailOrPhone(email, phone string) (bool, error)
	}

	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(tx *gorm.DB, user *entity.User) error {
	if tx == nil {
		tx = r.db
	}

	return tx.Create(user).Error
}

func (r *userRepository) ExistsByEmailOrPhone(email, phone string) (bool, error) {
	var count int64
	if err := r.db.Model(&entity.User{}).Where("email = ? OR phone = ?", email, phone).Count(&count).Error; err != nil {
		return false, err
	}
	return count > 0, nil
}
