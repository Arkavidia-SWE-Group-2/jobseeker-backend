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
		Detail(ctx *fiber.Ctx) error
		Update(ctx *fiber.Ctx) error
		Delete(ctx *fiber.Ctx) error
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

func (h *educationHandler) Detail(ctx *fiber.Ctx) error {
	educationID := ctx.Params("id")
	if educationID == "" {
		return response.NewFailed(domain.EDUCATION_DETAIL_FAILED, domain.ErrInvalidParameter, h.logger).Send(ctx)
	}

	res, err := h.usecase.Detail(ctx.Context(), educationID)
	if err != nil {
		return response.NewFailed(domain.EDUCATION_DETAIL_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.EDUCATION_DETAIL_SUCCESS, res, nil).Send(ctx)
}

func (h *educationHandler) Update(ctx *fiber.Ctx) error {
	educationID := ctx.Params("id")
	if educationID == "" {
		return response.NewFailed(domain.EDUCATION_UPDATE_FAILED, domain.ErrInvalidParameter, h.logger).Send(ctx)
	}

	var req domain.EducationUpdateRequest
	if err := h.validator.ParseAndValidate(ctx, &req); err != nil {
		return response.NewFailed(domain.EDUCATION_UPDATE_FAILED, err, h.logger).Send(ctx)
	}

	user, err := auth.ParseFromContext(ctx)
	if err != nil {
		return response.NewFailed(domain.EDUCATION_UPDATE_FAILED, err, h.logger).Send(ctx)
	}

	if err := h.usecase.Update(ctx.Context(), req, educationID, user.ID); err != nil {
		return response.NewFailed(domain.EDUCATION_UPDATE_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.EDUCATION_UPDATE_SUCCESS, nil, nil).Send(ctx)
}

func (h *educationHandler) Delete(ctx *fiber.Ctx) error {
	educationID := ctx.Params("id")
	if educationID == "" {
		return response.NewFailed(domain.EDUCATION_DELETE_FAILED, domain.ErrInvalidParameter, h.logger).Send(ctx)
	}

	user, err := auth.ParseFromContext(ctx)
	if err != nil {
		return response.NewFailed(domain.EDUCATION_DELETE_FAILED, err, h.logger).Send(ctx)
	}

	if err := h.usecase.Delete(ctx.Context(), educationID, user.ID); err != nil {
		return response.NewFailed(domain.EDUCATION_DELETE_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.EDUCATION_DELETE_SUCCESS, nil, nil).Send(ctx)
}
