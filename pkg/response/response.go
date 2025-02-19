package response

import (
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Success(c *fiber.Ctx, data interface{}) error {
	response := Response{
		Message: "success",
		Data: data,
	}
	return c.Status(fiber.StatusOK).JSON(response)
}
