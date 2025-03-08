package usecase

import (
	"github.com/gofiber/fiber/v2"
)

type (
	AuthUsecase interface {
		Verify(ctx *fiber.Ctx, id int) error
	}

	authUsecase struct {
	}
)

func NewAuthUsecase() AuthUsecase {
	return &authUsecase{}
}

func (u *authUsecase) Verify(ctx *fiber.Ctx, id int) error {
	return nil
}
