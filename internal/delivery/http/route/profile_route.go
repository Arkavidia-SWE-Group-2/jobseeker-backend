package route

import (
	"jobseeker/internal/delivery/http/handler"
	"jobseeker/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewProfileRoute(api *fiber.App, handler handler.ProfileHandler, middleware *middleware.Middleware) {
	router := api.Group("/profile")
	{
		router.Get("/:vanity", handler.GetProfile)
		router.Put("/", middleware.AuthMiddleware(), handler.UpdateProfile)
	}
}
