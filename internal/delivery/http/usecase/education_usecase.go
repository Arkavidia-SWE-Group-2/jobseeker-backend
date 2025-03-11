package usecase

import (
	"context"
	"jobseeker/internal/delivery/http/repository"
	"jobseeker/internal/domain"
	"jobseeker/internal/entity"
	"time"

	"gorm.io/gorm"
)

type (
	EducationUsecase interface {
		Create(ctx context.Context, req domain.EducationCreateRequest, userID string) error
	}

	educationUsecase struct {
		db            *gorm.DB
		educationRepo repository.EducationRepository
	}
)

func NewEducationUsecase(db *gorm.DB, educationRepo repository.EducationRepository) EducationUsecase {
	return &educationUsecase{db, educationRepo}
}

func (u *educationUsecase) Create(ctx context.Context, req domain.EducationCreateRequest, userID string) error {
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	startDate, err := time.Parse(time.DateOnly, req.StartDate)
	if err != nil {
		tx.Rollback()
		return err
	}
	endDate, err := time.Parse(time.DateOnly, req.EndDate)
	if err != nil {
		tx.Rollback()
		return err
	}

	education := entity.Education{
		UserID:      userID,
		School:      req.School,
		Degree:      req.Degree,
		Description: req.Description,
		StartDate:   startDate,
		EndDate:     endDate,
	}

	if err := u.educationRepo.Create(tx, &education); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
