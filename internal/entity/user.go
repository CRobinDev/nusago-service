package entity

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `json:"id" gorm:"type:uuid;primary_key;"`
	Fullname    string    `json:"fullname" gorm:"type:varchar(255);not null;"`
	Institution string    `json:"institution" gorm:"type:varchar(255);not null;"`
	Email       string    `json:"email" gorm:"type:varchar(255);unique;not null;"`
	Username    string    `json:"username" gorm:"type:varchar(255);unique;not null;"`
	Password    string    `json:"password" gorm:"type:varchar(255);not null;"`
	Description string 	  `json:"description" gorm:"type:text;"`
	CreatedAt   time.Time `json:"createdAt" gorm:"type:timestamp;autoCreateTime"`
	UpdatedAt   time.Time `json:"updatedAt" gorm:"type:timestamp;autoUpdateTime"`
}
