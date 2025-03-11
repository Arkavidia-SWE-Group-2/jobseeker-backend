package route

import (
	"jobseeker/internal/delivery/http/handler"
	"jobseeker/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewEducationRoute(api *fiber.App, handler handler.EducationHandler, middleware *middleware.Middleware) {
	router := api.Group("/educations")
	{
		router.Post("/", middleware.AuthMiddleware(), handler.Create)
	}
}
