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
		Detail(ctx context.Context, educationID string) (domain.EducationDetailResponse, error)
		Update(ctx context.Context, req domain.EducationUpdateRequest, educationID, userID string) error
		Delete(ctx context.Context, educationID, userID string) error
		GetAllByUser(ctx context.Context, userID string) ([]domain.EducationDetailResponse, error)
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

func (u *educationUsecase) Detail(ctx context.Context, educationID string) (domain.EducationDetailResponse, error) {
	tx := u.db.WithContext(ctx)

	res := domain.EducationDetailResponse{}

	var education entity.Education
	if err := u.educationRepo.FindByID(tx, educationID, &education); err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, domain.ErrEducationNotFound
		}
		return res, err
	}

	res.ID = education.ID
	res.School = education.School
	res.Degree = education.Degree
	res.Description = education.Description
	res.StartDate = education.StartDate.Format(time.DateOnly)
	res.EndDate = education.EndDate.Format(time.DateOnly)

	return res, nil
}

func (u *educationUsecase) Update(ctx context.Context, req domain.EducationUpdateRequest, educationID, userID string) error {
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	var education entity.Education
	if err := u.educationRepo.FindByIDAndUserID(tx, educationID, userID, &education); err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return domain.ErrEducationNotFound
		}
		return err
	}

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

	education.School = req.School
	education.Degree = req.Degree
	education.Description = req.Description
	education.StartDate = startDate
	education.EndDate = endDate

	if err := u.educationRepo.Update(tx, educationID, &education); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (u *educationUsecase) Delete(ctx context.Context, educationID, userID string) error {
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	exists, err := u.educationRepo.ExistsByIDAndUserID(tx, educationID, userID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if !exists {
		tx.Rollback()
		return domain.ErrEducationNotFound
	}

	if err := u.educationRepo.DeleteByID(tx, educationID); err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return domain.ErrEducationNotFound
		}
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (u *educationUsecase) GetAllByUser(ctx context.Context, userID string) ([]domain.EducationDetailResponse, error) {
	tx := u.db.WithContext(ctx)

	educations, err := u.educationRepo.GetAllByUserID(tx, userID)
	if err != nil {
		return nil, err
	}

	res := make([]domain.EducationDetailResponse, 0, len(educations))
	for _, education := range educations {
		res = append(res, domain.EducationDetailResponse{
			ID:          education.ID,
			School:      education.School,
			Degree:      education.Degree,
			Description: education.Description,
			StartDate:   education.StartDate.Format(time.DateOnly),
			EndDate:     education.EndDate.Format(time.DateOnly),
		})
	}

	return res, nil
}
