package middleware

import (
	"jobseeker/internal/domain"
	"jobseeker/internal/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AuthMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		token := ctx.Get("Authorization")

		if token == "" {
			return response.NewFailed("Unauthorized", fiber.NewError(fiber.StatusUnauthorized), nil).Send(ctx)
		}
		subs := strings.Split(token, "Bearer ")
		if len(subs) != 2 {
			return response.NewFailed("Unauthorized", fiber.NewError(fiber.StatusUnauthorized), nil).Send(ctx)
		}

		claims, err := m.jwt.VerifyToken(subs[1])
		if err != nil {
			m.log.Error(err)
			return response.NewFailed("Unauthorized", fiber.NewError(fiber.StatusUnauthorized), nil).Send(ctx)
		}

		user, err := m.authUsecase.Verify(ctx.Context(), claims.ID)
		if err != nil {
			m.log.Error(err)
			return response.NewFailed("Unauthorized", fiber.NewError(fiber.StatusUnauthorized), nil).Send(ctx)
		}

		ctx.Locals(domain.AUTH_USER, user)
		return ctx.Next()
	}
}
