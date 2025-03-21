package repository

import (
	"jobseeker/internal/entity"

	"gorm.io/gorm"
)

type (
	ProfileRepository interface {
		GetProfileByVanity(tx *gorm.DB, vanity string, profile *entity.Profile) error
	}

	profileRepository struct {
		db *gorm.DB
	}
)

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}

func (r *profileRepository) GetProfileByVanity(tx *gorm.DB, vanity string, profile *entity.Profile) error {
	if tx == nil {
		tx = r.db
	}
	query := tx.
		Joins("JOIN users ON users.id = profiles.user_id").
		Where("users.vanity = ?", vanity).
		Select("profiles.*")

	if err := query.First(profile).Error; err != nil {
		return err
	}

	return nil
}
