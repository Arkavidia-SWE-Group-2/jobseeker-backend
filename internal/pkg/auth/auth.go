package auth

import (
	"jobseeker/internal/domain"
	"jobseeker/internal/entity"

	"github.com/gofiber/fiber/v2"
)

func ParseFromContext(ctx *fiber.Ctx) (entity.User, error) {
	user, ok := ctx.Locals(domain.AUTH_USER).(entity.User)
	if !ok {
		return entity.User{}, fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	}
	return user, nil
}
