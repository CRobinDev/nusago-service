package service

import (
	"time"

	"github.com/CRobinDev/nusago-service/internal/dto"
	"github.com/CRobinDev/nusago-service/internal/entity"
	"github.com/CRobinDev/nusago-service/internal/repository"
	"github.com/CRobinDev/nusago-service/pkg/gomail"
	"github.com/CRobinDev/nusago-service/pkg/helper"
	"github.com/CRobinDev/nusago-service/pkg/jwt"
	"github.com/CRobinDev/nusago-service/pkg/response"
	"github.com/CRobinDev/nusago-service/pkg/validator"
	"github.com/google/uuid"
)

type IUserService interface {
	Register(req dto.RegisterRequest) error
	Login(req dto.LoginRequest) (dto.LoginResponse, error)
	GetUser(req dto.TokenLoginRequest) (dto.TokenLoginResponse, error)
	Update(req dto.UpdateRequest) error
	Delete(req dto.DeleteRequest) error
	SendNotification(req dto.NotificationRequest) error
}

type userService struct {
	ur     repository.IUserRepository
	jwt    jwt.IJWT
	gomail *gomail.Gomail
}

func NewUserService(ur repository.IUserRepository, jwt jwt.IJWT, gomail *gomail.Gomail) IUserService {
	return &userService{
		ur:     ur,
		jwt:    jwt,
		gomail: gomail,
	}
}
func (us *userService) Register(req dto.RegisterRequest) error {
	if err := ValidateRequestRegister(req); err != nil {
		return err
	}

	hashedPassword, err := helper.HashPassword(req.Password)
	if err != nil {
		return &response.ErrHashPassword
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	createdAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	user := entity.User{
		ID:          uuid.New(),
		Fullname:    req.Fullname,
		Institution: req.Institution,
		Username:    req.Username,
		Email:       req.Email,
		Password:    hashedPassword,
		CreatedAt:   createdAt,
		UpdatedAt:   createdAt,
	}

	if err := us.ur.Create(&user); err != nil {
		return err
	}

	return nil
}

func (us *userService) Login(req dto.LoginRequest) (dto.LoginResponse, error) {
	user, err := us.ur.FindByEmail(req.Email)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	if err := helper.ComparePassword(user.Password, req.Password); err != nil {
		return dto.LoginResponse{}, err
	}

	token, err := us.jwt.CreateToken(&user)
	if err != nil {
		return dto.LoginResponse{}, err
	}

	resp := dto.LoginResponse{
		Username: user.Username,
		ID:       user.ID,
		Token:    token,
	}

	return resp, nil
}

func (us *userService) GetUser(req dto.TokenLoginRequest) (dto.TokenLoginResponse, error) {
	user, err := us.ur.FindByID(req.ID)
	if err != nil {
		return dto.TokenLoginResponse{}, err
	}

	return dto.TokenLoginResponse{
		ID:          user.ID,
		Username:    user.Username,
		Fullname:    user.Fullname,
		Institution: user.Institution,
	}, nil
}

func (us *userService) Update(req dto.UpdateRequest) error {
	user, err := us.ur.FindByID(req.ID)
	if err != nil {
		return err
	}

	if req.Fullname != "" {
		user.Fullname = req.Fullname
	}

	if req.Institution != "" {
		user.Institution = req.Institution
	}

	if req.Description != "" {
		user.Description = req.Description
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05", time.Now().In(loc).Format("2006-01-02 15:04:05"))

	user.UpdatedAt = updatedAt

	if err := us.ur.Update(&user); err != nil {
		return err
	}

	return nil
}

func (us *userService) Delete(req dto.DeleteRequest) error {
	return us.ur.Delete(req.ID)
}

func (us *userService) SendNotification(req dto.NotificationRequest) error {
	user, err := us.ur.FindByID(req.ID)
	if err != nil {
		return err
	}
	req.Email = user.Email
	if err := us.gomail.SendNotification(req); err != nil {
		return &response.ErrFailedSendNotification
	}

	return nil
}

func ValidateRequestRegister(req dto.RegisterRequest) error {
	switch {
	case req.Email == "" || !validator.ValidateEmail(req.Email):
		return &response.ErrInvalidEmail
	case req.Password == "" || !validator.ValidatePassword(req.Password):
		return &response.ErrInvalidPassword
	}
	return nil
}
