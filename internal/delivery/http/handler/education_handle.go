package handler

import (
	"jobseeker/internal/delivery/http/usecase"
	"jobseeker/internal/domain"
	"jobseeker/internal/pkg/auth"
	"jobseeker/internal/pkg/response"
	"jobseeker/internal/pkg/validate"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
)

type (
	EducationHandler interface {
		Create(ctx *fiber.Ctx) error
	}

	educationHandler struct {
		validator *validate.Validator
		logger    *logrus.Logger
		usecase   usecase.EducationUsecase
	}
)

func NewEducationHandler(validator *validate.Validator, logger *logrus.Logger, usecase usecase.EducationUsecase) EducationHandler {
	return &educationHandler{validator, logger, usecase}
}

func (h *educationHandler) Create(ctx *fiber.Ctx) error {
	var req domain.EducationCreateRequest
	if err := h.validator.ParseAndValidate(ctx, &req); err != nil {
		return response.NewFailed(domain.EDUCATION_CREATE_FAILED, err, h.logger).Send(ctx)
	}

	user, err := auth.ParseFromContext(ctx)
	if err != nil {
		return response.NewFailed(domain.EDUCATION_CREATE_FAILED, err, h.logger).Send(ctx)
	}

	if err := h.usecase.Create(ctx.Context(), req, user.ID); err != nil {
		return response.NewFailed(domain.EDUCATION_CREATE_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.EDUCATION_CREATE_SUCCESS, nil, nil).Send(ctx)
}
