package repository

import "gorm.io/gorm"

type (
	ProfileRepository interface{}

	profileRepository struct {
		db *gorm.DB
	}
)

func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db}
}
