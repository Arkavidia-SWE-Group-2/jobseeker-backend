package bootstrap

import (
	"jobseeker/internal/delivery/http/handler"
	"jobseeker/internal/delivery/http/middleware"
	"jobseeker/internal/delivery/http/repository"
	"jobseeker/internal/delivery/http/route"
	"jobseeker/internal/delivery/http/usecase"
	"jobseeker/internal/pkg/validate"
	"jobseeker/pkg/jwt"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

type ApiBootstrapConfig struct {
	Api       *fiber.App
	Config    *viper.Viper
	DB        *gorm.DB
	Log       *logrus.Logger
	Validator *validate.Validator
	JWT       *jwt.JWT
}

func ApiBootstrap(conf *ApiBootstrapConfig) {
	/**--------------------------------------------
	 **  REPOSITORIES
	 *---------------------------------------------**/
	userRepo := repository.NewUserRepository(conf.DB)
	profileRepo := repository.NewProfileRepository(conf.DB)

	/**--------------------------------------------
	 **  USECASES
	 *---------------------------------------------**/
	authUsecase := usecase.NewAuthUsecase(conf.DB, conf.JWT, userRepo, profileRepo)

	/**--------------------------------------------
	**  HANDLERS
	*---------------------------------------------**/
	baseHandler := handler.NewBaseHandler()
	authHandler := handler.NewAuthHandler(conf.Validator, conf.Log, authUsecase)

	/**--------------------------------------------
	**  MIDDLEWARE & ROUTE SETUP
	*---------------------------------------------**/
	middleware := middleware.NewMiddleware(conf.Log, conf.JWT, conf.Config, authUsecase)
	route.Setup(&route.RouteConfig{
		Api:         conf.Api,
		Middleware:  middleware,
		BaseHandler: baseHandler,
		AuthHandler: authHandler,
	})
}
