package handler

import (
	"jobseeker/internal/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type (
	BaseHandler interface {
		HelloWorld(ctx *fiber.Ctx) error
	}

	baseHandler struct {
	}
)

func NewBaseHandler() BaseHandler {
	return &baseHandler{}
}

func (h *baseHandler) HelloWorld(ctx *fiber.Ctx) error {
	return response.NewSuccess("Hello World", nil, nil).Send(ctx)
}
