package usecase

import (
	"jobseeker/internal/delivery/http/repository"

	"gorm.io/gorm"
)

type (
	EducationUsecase interface{}

	educationUsecase struct {
		db            *gorm.DB
		educationRepo repository.EducationRepository
	}
)

func NewEducationUsecase(db *gorm.DB, educationRepo repository.EducationRepository) EducationUsecase {
	return &educationUsecase{db, educationRepo}
}
