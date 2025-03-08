package route

import (
	"jobseeker/internal/delivery/http/handler"

	"github.com/gofiber/fiber/v2"
)

func NewBaseRoute(api *fiber.App, handler handler.BaseHandler) {
	api.Get("/", handler.HelloWorld)
}
