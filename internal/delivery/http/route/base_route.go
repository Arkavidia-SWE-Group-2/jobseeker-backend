package route

import (
	"jobseeker/internal/delivery/http/handler"
	"jobseeker/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
)

func NewBaseRoute(api *fiber.App, handler handler.BaseHandler, middleware *middleware.Middleware) {
	api.Get("/", handler.HelloWorld)
	api.Get("/testauth", middleware.AuthMiddleware(), handler.HelloWorld)
}
