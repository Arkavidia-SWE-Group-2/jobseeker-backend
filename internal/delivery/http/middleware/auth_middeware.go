package middleware

import (
	"jobseeker/internal/domain"
	"jobseeker/internal/pkg/response"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (m *Middleware) AuthMiddleware(roles ...string) fiber.Handler {
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
			return response.NewFailed("Unauthorized", err, nil).Send(ctx)
		}

		err = m.authUsecase.Verify(ctx.Context(), claims.ID)
		if err != nil {
			return response.NewFailed("Unauthorized", err, nil).Send(ctx)
		}

		if len(roles) > 0 {
			role := claims.Role
			if role == "" {
				return response.NewFailed("Forbidden", fiber.NewError(fiber.StatusUnauthorized), nil).Send(ctx)
			}
			isAccepted := false
			for _, r := range roles {
				if r == role {
					isAccepted = true
					break
				}
			}
			if !isAccepted {
				return response.NewFailed("Forbidden", fiber.NewError(fiber.StatusForbidden), nil).Send(ctx)
			}
		}

		ctx.Locals(domain.AUTH_USER, claims)
		return ctx.Next()
	}
}
