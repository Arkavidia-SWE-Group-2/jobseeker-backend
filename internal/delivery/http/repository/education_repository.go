package repository

import "gorm.io/gorm"

type (
	EducationRepository interface{}

	educationRepository struct {
		db *gorm.DB
	}
)

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepository{db}
}
