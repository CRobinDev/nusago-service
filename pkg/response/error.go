package response

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

type Errors struct {
	Code int
	Err  error
}

func (e *Errors) Error() string {
	return e.Err.Error()
}

func NewError(code int, err string) Errors {
	return Errors{
		Code: code,
		Err:  errors.New(err),
	}
}

var (

	//Database
	ErrConnectDatabase = NewError(fiber.StatusInternalServerError, "failed to connect database")

	ErrMigrateDatabase = NewError(fiber.StatusInternalServerError, "failed migrate database")

	//User
	ErrUserNotFound = NewError(fiber.StatusNotFound, "user not found")

	ErrUserAlreadyExists = NewError(fiber.StatusConflict, "user already exists")

	ErrHashPassword = NewError(fiber.StatusInternalServerError, "something went wrong!")

	ErrGenerateToken = NewError(fiber.StatusInternalServerError, "an unexpected error occured")

	ErrInvalidEmail = NewError(fiber.StatusBadRequest, "invalid email")

	ErrInvalidPassword = NewError(fiber.StatusBadRequest, "invalid password")

	ErrJWTToken = NewError(fiber.StatusInternalServerError, "something went wrong!")

	ErrFailedSendNotification = NewError(fiber.StatusInternalServerError, "failed to notification")

	ErrUnauthorized = NewError(fiber.StatusUnauthorized, "invalid login credentials")

	ErrSetHTML = NewError(fiber.StatusInternalServerError, "failed to set HTML template")
	
	ErrExecuteHTML = NewError(fiber.StatusInternalServerError, "failed to set HTML template")

)
