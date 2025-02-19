package middleware

import (
	"strings"

	"github.com/CRobinDev/nusago-service/pkg/jwt"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func Authenticate(jwtService jwt.IJWT) fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if bearer == "" {
			return &response.ErrUnauthorized
		}
		tokenSlice := strings.Split(bearer, " ")
		if len(tokenSlice) != 2 {
			return &response.ErrUnauthorized
		}

		token := tokenSlice[1]
		userID, err := jwtService.DecodeToken(token)

		if err != nil {
			return &response.ErrUnauthorized
		}

		c.Locals("claims", userID)

		return c.Next()
	}
}

func AuthenticateGRPC() fiber.Handler {
	return func(c *fiber.Ctx) error {
		bearer := c.Get("Authorization")
		if bearer == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		tokenSlice := strings.Split(bearer, " ")
		if len(tokenSlice) != 2 {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		if tokenSlice[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Unauthorized"})
		}

		c.Locals("token", tokenSlice[1])

		return c.Next()
	}
}
