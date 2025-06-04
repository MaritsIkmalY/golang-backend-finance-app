package models

import (
	"time"
)

type UserResponse struct {
	ID        uint      `json:"id"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Token     string    `json:"token,omitempty"`
}

type RegisterUserRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Name            string `json:"name" validate:"required"`
	Password        string `json:"password" validate:"required,min=6"`
	PasswordConfirm string `json:"password_confirm" validate:"required,eqfield=Password"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}
