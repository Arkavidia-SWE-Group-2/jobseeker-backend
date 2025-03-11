package route

import (
	"jobseeker/internal/delivery/http/handler"
	"jobseeker/internal/delivery/http/middleware"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type RouteConfig struct {
	Api              *fiber.App
	Middleware       *middleware.Middleware
	BaseHandler      handler.BaseHandler
	AuthHandler      handler.AuthHandler
	EducationHandler handler.EducationHandler
}

func Setup(c *RouteConfig) {
	c.Api.Use(recover.New())
	c.Api.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	c.Api.Use(c.Middleware.CorsMiddleware())

	NewBaseRoute(c.Api, c.BaseHandler, c.Middleware)
	NewAuthRoute(c.Api, c.AuthHandler)
	NewEducationRoute(c.Api, c.EducationHandler, c.Middleware)
}
