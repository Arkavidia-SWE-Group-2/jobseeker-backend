package handler

import (
	"jobseeker/internal/delivery/http/usecase"
	"jobseeker/internal/domain"
	"jobseeker/internal/pkg/response"
	"jobseeker/internal/pkg/validate"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	AuthHandler interface {
		Login(ctx *fiber.Ctx) error
		Register(ctx *fiber.Ctx) error
	}

	authHandler struct {
		validator *validate.Validator
		logger    *logrus.Logger
		usecase   usecase.AuthUsecase
	}
)

func NewAuthHandler(validator *validate.Validator, logger *logrus.Logger, usecase usecase.AuthUsecase) AuthHandler {
	return &authHandler{validator, logger, usecase}
}

func (h *authHandler) Login(ctx *fiber.Ctx) error {
	panic("unimplemented")
}

func (h *authHandler) Register(ctx *fiber.Ctx) error {
	req := domain.AuthRegisterRequest{}

	if err := h.validator.ParseAndValidate(ctx, &req); err != nil {
		return response.NewFailed(domain.AUTH_REGISTER_FAILED, err, h.logger).Send(ctx)
	}

	if err := h.usecase.Register(ctx.Context(), req); err != nil {
		return response.NewFailed(domain.AUTH_REGISTER_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.AUTH_REGISTER_SUCCESS, nil, nil).Send(ctx)
}
