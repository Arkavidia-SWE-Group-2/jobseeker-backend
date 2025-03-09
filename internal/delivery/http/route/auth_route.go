package route

import (
	"jobseeker/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewAuthRoute(api *fiber.App, handler handler.AuthHandler) {
	router := api.Group("/auth")
	{
		router.Post("/register", handler.Register)
		router.Post("/login", handler.Login)
	}
}
