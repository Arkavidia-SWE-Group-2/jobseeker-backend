package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func (m *Middleware) CorsMiddleware() fiber.Handler {
	return cors.New(cors.Config{
		AllowHeaders:  "Origin, Content-Type, Accept, Authorization, Content-Length, Accept-Encoding",
		AllowMethods:  "GET, POST, PUT, PATCH, DELETE",
		AllowOrigins:  m.config.GetString("api.cors.origins"),
		ExposeHeaders: "Content-Length, Content-Type",
	})
}
