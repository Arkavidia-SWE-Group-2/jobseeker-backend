package usecase

import (
	"context"
	"jobseeker/internal/delivery/http/repository"
	"jobseeker/internal/domain"
	"jobseeker/internal/entity"

	"gorm.io/gorm"
)

type (
	ProfileUsecase interface {
		GetProfile(ctx context.Context, vanity string) (domain.ProfileResponse, error)
		UpdateProfile(ctx context.Context, req domain.ProfileEditRequest, userID string) error
	}

	profileUsecase struct {
		db          *gorm.DB
		profileRepo repository.ProfileRepository
	}
)

func NewProfileUsecase(db *gorm.DB, profileRepo repository.ProfileRepository) ProfileUsecase {
	return &profileUsecase{db, profileRepo}
}

func (p *profileUsecase) GetProfile(ctx context.Context, vanity string) (domain.ProfileResponse, error) {
	var profile entity.Profile

	res := domain.ProfileResponse{}
	if err := p.profileRepo.GetProfileByVanity(p.db, vanity, &profile); err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, domain.ErrProfileNotFound
		}
		return res, err
	}

	res.ID = profile.UserID
	res.FirstName = profile.FirstName
	res.LastName = profile.LastName
	res.About = profile.About
	res.Headline = profile.Headline
	res.Photo = profile.Photo
	res.Vanity = vanity

	return res, nil
}

func (p *profileUsecase) UpdateProfile(ctx context.Context, req domain.ProfileEditRequest, userID string) error {
	var profile entity.Profile

	profile.UserID = userID
	profile.FirstName = req.FirstName
	profile.LastName = req.LastName
	profile.Headline = req.Headline
	profile.About = req.About

	if err := p.profileRepo.UpdateProfile(p.db, &profile); err != nil {
		return err
	}

	return nil
}
