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
