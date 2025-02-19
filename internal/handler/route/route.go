package route

import (
	"github.com/CRobinDev/nusago-service/internal/handler"
	"github.com/CRobinDev/nusago-service/internal/middleware"
	"github.com/CRobinDev/nusago-service/pkg/jwt"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	App         *fiber.App
	UserHandler handler.IUserHandler
	Jwt         jwt.IJWT
}

func (c *Config) Register() {
	api := c.App.Group("/api/v1")

	api.Get("/health-check", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	c.userRoutes(api)
}

func (c *Config) userRoutes(r fiber.Router) {
	user := r.Group("/users")
	user.Post("/register", c.UserHandler.Register())
	user.Post("/login", c.UserHandler.Login())
	user.Get("/me", middleware.Authenticate(c.Jwt), c.UserHandler.GetUser())
	user.Patch("/update-account", middleware.Authenticate(c.Jwt), c.UserHandler.Update())
	user.Delete("/delete-account", middleware.Authenticate(c.Jwt), c.UserHandler.Delete())
	user.Post("/notification", middleware.Authenticate(c.Jwt), c.UserHandler.Notification())
}

