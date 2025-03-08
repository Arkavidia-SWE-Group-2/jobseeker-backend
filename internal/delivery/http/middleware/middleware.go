package middleware

import (
	"jobseeker/internal/delivery/http/usecase"
	"jobseeker/pkg/jwt"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type Middleware struct {
	log         *logrus.Logger
	jwt         *jwt.JWT
	config      *viper.Viper
	authUsecase usecase.AuthUsecase
}

func NewMiddleware(log *logrus.Logger, jwt *jwt.JWT, config *viper.Viper, authUsecase usecase.AuthUsecase) *Middleware {
	return &Middleware{log, jwt, config, authUsecase}
}
