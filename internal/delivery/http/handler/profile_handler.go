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
	ProfileHandler interface {
		GetProfile(ctx *fiber.Ctx) error
		UpdateProfile(ctx *fiber.Ctx) error
	}

	profileHandler struct {
		validator *validate.Validator
		logger    *logrus.Logger
		usecase   usecase.ProfileUsecase
	}
)

func NewProfileHandler(validator *validate.Validator, logger *logrus.Logger, usecase usecase.ProfileUsecase) ProfileHandler {
	return &profileHandler{validator, logger, usecase}
}

func (h *profileHandler) GetProfile(ctx *fiber.Ctx) error {
	vanity := ctx.Params("vanity")

	res, err := h.usecase.GetProfile(ctx.Context(), vanity)
	if err != nil {
		return response.NewFailed(domain.PROFILE_GET_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.PROFILE_GET_SUCCESS, res, nil).Send(ctx)
}

func (h *profileHandler) UpdateProfile(ctx *fiber.Ctx) error {
	user, err := auth.ParseFromContext(ctx)
	if err != nil {
		return response.NewFailed(domain.PROFILE_GET_FAILED, err, h.logger).Send(ctx)
	}

	var req domain.ProfileEditRequest
	if err := h.validator.ParseAndValidate(ctx, &req); err != nil {
		return response.NewFailed(domain.PROFILE_UPDATE_FAILED, err, h.logger).Send(ctx)
	}

	if err := h.usecase.UpdateProfile(ctx.Context(), req, user.ID); err != nil {
		return response.NewFailed(domain.PROFILE_UPDATE_FAILED, err, h.logger).Send(ctx)
	}

	return response.NewSuccess(domain.PROFILE_UPDATE_SUCCESS, nil, nil).Send(ctx)
}
