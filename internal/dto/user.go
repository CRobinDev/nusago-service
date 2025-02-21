package dto

import (
	"github.com/google/uuid"
)

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=24"`
}

type RegisterRequest struct {
	Fullname    string `json:"fullname" validate:"required,min=5,max=255"`
	Institution string `json:"institution" validate:"required,min=5,max=255"`
	Email       string `json:"email" validate:"required,email"`
	Username    string `json:"username" validate:"required,min=5,max=64"`
	Password    string `json:"password" validate:"required,min=8,max=24"`
}

type TokenLoginRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type UpdateRequest struct {
	ID          uuid.UUID `json:"-"`
	Fullname    string    `json:"fullname" validate:"required,min=5,max=255"`
	Institution string    `json:"institution" validate:"required,min=5,max=255"`
	Description string    `json:"description" validate:"required,min=8"`
}

type DeleteRequest struct {
	ID uuid.UUID `json:"id" validate:"required"`
}

type NotificationRequest struct {
	ID       uuid.UUID `json:"-"`
	Email    string    `json:"-"`
	Fullname string    `json:"fullname" validate:"required"`
	Feature  string    `json:"feature" validate:"required,oneof=Blog Portofolio"`
	Link     string    `json:"link" validate:"required"`
}

type LoginResponse struct {
	Username string    `json:"username"`
	ID       uuid.UUID `json:"userID"`
	Token    string    `json:"token"`
}

type TokenLoginResponse struct {
	ID          uuid.UUID `json:"userID"`
	Username    string    `json:"username"`
	Fullname    string    `json:"fullname"`
	Institution string    `json:"institution"`
	Description string    `json:"description"`
}

