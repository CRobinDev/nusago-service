package config

import (
	"github.com/CRobinDev/nusago-service/internal/handler"
	"github.com/CRobinDev/nusago-service/internal/handler/route"
	"github.com/CRobinDev/nusago-service/internal/repository"
	"github.com/CRobinDev/nusago-service/internal/service"
	"github.com/CRobinDev/nusago-service/pkg/gomail"
	"github.com/CRobinDev/nusago-service/pkg/jwt"
	"github.com/CRobinDev/nusago-service/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type AppConfig struct {
	App  *fiber.App
	DB   *gorm.DB
}

func StartApp(config *AppConfig) {
	jwt := jwt.Init()
	val := validator.NewValidator()
	gomail := gomail.NewGomail()

	userRepository := repository.NewUserRepository(config.DB)
	userService := service.NewUserService(userRepository, jwt, gomail)
	userHandler := handler.NewUserHandler(userService, val)
	routes := route.Config{
		App:         config.App,
		UserHandler: userHandler,
		Jwt: jwt,
	}

	routes.Register()
}
