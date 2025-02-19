package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors() func(*fiber.Ctx) error {
	return cors.New(cors.Config{
		AllowOrigins:  "http://localhost:3000",
		AllowMethods:  "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS", 
		AllowHeaders:  "Access-Control-Allow-Headers, Origin, Accept, Content-Type, Authorization, X-CSRF-Token",
		AllowCredentials: true,
	})
}
