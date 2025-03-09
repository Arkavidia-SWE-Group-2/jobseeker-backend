package usecase

import (
	"context"
	"jobseeker/internal/delivery/http/repository"
	"jobseeker/internal/domain"
	"jobseeker/internal/entity"
	"jobseeker/internal/pkg/helper"
	"jobseeker/pkg/jwt"

	"gorm.io/gorm"
)

type (
	AuthUsecase interface {
		Verify(ctx context.Context, id string) error
		Login(ctx context.Context, req domain.AuthLoginRequest) (domain.AuthLoginResponse, error)
		Register(ctx context.Context, req domain.AuthRegisterRequest) error
	}

	authUsecase struct {
		db          *gorm.DB
		jwt         *jwt.JWT
		userRepo    repository.UserRepository
		profileRepo repository.ProfileRepository
	}
)

func NewAuthUsecase(db *gorm.DB, jwt *jwt.JWT, userRepo repository.UserRepository, profileRepo repository.ProfileRepository) AuthUsecase {
	return &authUsecase{db, jwt, userRepo, profileRepo}
}

func (u *authUsecase) Verify(ctx context.Context, id string) error {
	return nil
}

func (u *authUsecase) Login(ctx context.Context, req domain.AuthLoginRequest) (domain.AuthLoginResponse, error) {
	res := domain.AuthLoginResponse{}

	var user entity.User
	if err := u.userRepo.GetByCredential(&user, req.Credential); err != nil {
		if err == gorm.ErrRecordNotFound {
			return res, domain.ErrWrongCredential
		}
		return res, err
	}

	if err := helper.ComparePassword(req.Password, user.Password); err != nil {
		return res, domain.ErrWrongCredential
	}

	token, err := u.jwt.GenerateToken(jwt.Payload{Sub: user.ID.String()})
	if err != nil {
		return res, err
	}
	res.Token = token

	return res, nil
}

func (u *authUsecase) Register(ctx context.Context, req domain.AuthRegisterRequest) error {
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			panic(r)
		}
	}()

	exists, err := u.userRepo.ExistsByEmailOrPhone(req.Email, req.Phone)
	if err != nil {
		tx.Rollback()
		return err
	}
	if exists {
		tx.Rollback()
		return domain.ErrEmailOrPhoneAlreadyExists
	}

	vanity := helper.VanityFromEmail(req.Email)

	user := entity.User{
		Email:    req.Email,
		Password: req.Password,
		Phone:    req.Phone,
		Vanity:   vanity,
		Profile: &entity.Profile{
			FirstName: req.FirstName,
			LastName:  req.LastName,
		},
	}

	if err := u.userRepo.Create(tx, &user); err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
