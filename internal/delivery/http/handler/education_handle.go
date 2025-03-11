package handler

import (
	"jobseeker/internal/delivery/http/usecase"
	"jobseeker/internal/pkg/validate"

	"github.com/sirupsen/logrus"
)

type (
	EducationHandler interface{}

	educationHandler struct {
		validator *validate.Validator
		logger    *logrus.Logger
		usecase   usecase.EducationUsecase
	}
)

func NewEducationHandler(validator *validate.Validator, logger *logrus.Logger, usecase usecase.EducationUsecase) EducationHandler {
	return &educationHandler{validator, logger, usecase}
}
