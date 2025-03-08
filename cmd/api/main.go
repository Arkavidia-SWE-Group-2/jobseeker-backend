package main

import (
	"fmt"
	"jobseeker/database"
	"jobseeker/internal/bootstrap"
	"jobseeker/internal/config"
	"jobseeker/internal/pkg/validate"
	"jobseeker/pkg/jwt"
)

func main() {
	viperConfig := config.NewViper()
	log := config.NewLogger(viperConfig)
	db := database.New(viperConfig)
	validator := validate.NewValidator()
	api := config.NewAPI(viperConfig, log)
	jwt := jwt.New()

	bootstrap.ApiBootstrap(&bootstrap.ApiBootstrapConfig{
		Config:    viperConfig,
		Log:       log,
		DB:        db,
		Validator: validator,
		Api:       api,
		JWT:       jwt,
	})

	listeningHost := viperConfig.GetString("api.host")
	listeningPort := viperConfig.GetInt("api.port")

	if err := api.Listen(fmt.Sprintf("%s:%d", listeningHost, listeningPort)); err != nil {
		panic(fmt.Errorf("fatal error failed to start api server: %w", err))
	}
}
