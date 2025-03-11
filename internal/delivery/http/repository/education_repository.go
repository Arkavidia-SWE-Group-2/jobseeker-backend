package repository

import (
	"jobseeker/internal/entity"

	"gorm.io/gorm"
)

type (
	EducationRepository interface {
		Create(tx *gorm.DB, education *entity.Education) error
	}

	educationRepository struct {
		db *gorm.DB
	}
)

func NewEducationRepository(db *gorm.DB) EducationRepository {
	return &educationRepository{db}
}

func (r *educationRepository) Create(tx *gorm.DB, education *entity.Education) error {
	if tx == nil {
		tx = r.db
	}
	return tx.Create(education).Error
}
