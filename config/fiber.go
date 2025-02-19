package config

import (
	"encoding/json"
	"errors"

	"github.com/CRobinDev/nusago-service/internal/middleware"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"github.com/CRobinDev/nusago-service/pkg/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
)

func NewFiber() *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: newErrorHandler(),
	})

	app.Use(func(c *fiber.Ctx) error {
		if c.Method() == fiber.MethodOptions {
			return c.SendStatus(fiber.StatusNoContent) 
		}
		return c.Next()
	})
	
	app.Use(middleware.Cors())

	return app
}

func newErrorHandler() fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		var apiError *response.Errors
		if errors.As(err, &apiError) {
			return c.Status(apiError.Code).JSON(fiber.Map{
				"errors": fiber.Map{"message": apiError.Error()},
			})
		}

		var fiberError *fiber.Error
		if errors.As(err, &fiberError) {
			return c.Status(fiberError.Code).JSON(fiber.Map{
				"errors": fiber.Map{
					"message": utils.StatusMessage(fiberError.Code),
					"err":     err,
				},
			})
		}

		var validationError validator.ValidationErrors
		if errors.As(err, &validationError) {
			validationDetails := fiber.Map{}
			for field, msg := range validationError {
				validationDetails[field] = msg
			}

			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"errors": validationDetails,
			})
		}

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"errors": fiber.Map{
				"message": utils.StatusMessage(fiber.StatusInternalServerError),
				"err":     err,
			},
		})
	}
}
