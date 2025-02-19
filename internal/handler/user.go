package handler

import (
	"github.com/CRobinDev/nusago-service/internal/dto"
	"github.com/CRobinDev/nusago-service/internal/service"
	"github.com/CRobinDev/nusago-service/pkg/jwt"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"github.com/CRobinDev/nusago-service/pkg/validator"
	"github.com/gofiber/fiber/v2"
)

type IUserHandler interface {
	Register() fiber.Handler
	Login() fiber.Handler
	GetUser() fiber.Handler
	Update() fiber.Handler
	Delete() fiber.Handler
	Notification() fiber.Handler
}

type userHandler struct {
	us  service.IUserService
	val validator.Validator
}

func NewUserHandler(us service.IUserService, val validator.Validator) IUserHandler {
	return &userHandler{
		us:  us,
		val: val,
	}
}

func (uh *userHandler) Register() fiber.Handler {
	return func(c *fiber.Ctx) error {
		req := dto.RegisterRequest{}
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}

		valErr := uh.val.Validate(req)
		if valErr != nil {
			return valErr
		}

		err = uh.us.Register(req)
		if err != nil {
			return err
		}

		return response.Success(c, nil)
	}
}

func (uh *userHandler) Login() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.LoginRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}

		valErr := uh.val.Validate(req)
		if valErr != nil {
			return valErr
		}

		resp, err := uh.us.Login(req)
		if err != nil {
			return err
		}

		return response.Success(c, resp)
	}
}

func (uh *userHandler) Delete() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.DeleteRequest
		userID, err := jwt.GetUser(c)
		if err != nil {
			return err
		}

		req.ID = userID
		valErr := uh.val.Validate(req)
		if valErr != nil {
			return valErr
		}

		err = uh.us.Delete(req)
		if err != nil {
			return err
		}

		return response.Success(c, nil)
	}
}

func (uh *userHandler) GetUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID, err := jwt.GetUser(c)
		if err != nil {
			return err
		}

		req := dto.TokenLoginRequest{
			ID: userID,
		}

		resp, err := uh.us.GetUser(req)
		if err != nil {
			return err
		}

		return response.Success(c, resp)
	}
}

func (uh *userHandler) Update() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.UpdateRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}

		userID, err := jwt.GetUser(c)
		if err != nil {
			return err
		}

		req.ID = userID
		valErr := uh.val.Validate(req)
		if valErr != nil {
			return valErr
		}

		err = uh.us.Update(req)
		if err != nil {
			return err
		}

		return response.Success(c, nil)

	}
}

func (uh *userHandler) Notification() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var req dto.NotificationRequest
		err := c.BodyParser(&req)
		if err != nil {
			return err
		}
		userID, err := jwt.GetUser(c)
		if err != nil {
			return err
		}

		req.ID = userID
		valErr := uh.val.Validate(req)
		if valErr != nil {
			return valErr
		}

		err = uh.us.SendNotification(req)
		if err != nil {
			return err
		}

		return response.Success(c, nil)
	}
}

// func parseFeatureRequest(feature string) (dto.Features, error) {
// 	switch feature {
// 	case "portofolio":
// 		return dto.Portofolio, nil
// 	case "blog":
// 		return dto.Blog, nil
// 	default:
// 		return dto.Unknown, errors.New("invalid feature")
// 	}
// }
