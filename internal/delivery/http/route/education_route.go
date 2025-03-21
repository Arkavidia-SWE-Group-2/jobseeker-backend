package route

import (
	"jobseeker/internal/delivery/http/handler"
	"jobseeker/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewEducationRoute(api *fiber.App, handler handler.EducationHandler, middleware *middleware.Middleware) {
	router := api.Group("/educations")
	{
		router.Get("/", middleware.AuthMiddleware(), handler.GetAllByUser)
		router.Post("/", middleware.AuthMiddleware(), handler.Create)
		router.Get("/:id", handler.Detail)
		router.Put("/:id", middleware.AuthMiddleware(), handler.Update)
		router.Delete("/:id", middleware.AuthMiddleware(), handler.Delete)
	}
}
