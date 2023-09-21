package model

import (
	"time"

	"github.com/google/uuid"
)

type ResponseBody[T any] struct {
	Status  string  `json:"status"`
	Message *string `json:"message"`
	Data    *T      `json:"data"`
}

type AppResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserResponse struct {
	ID                         string `json:"id"`
	AppID                      string `json:"app_id"`
	DisplayName                string `json:"display_name"`
	FirstName                  string `json:"first_name"`
	LastName                   string `json:"last_name"`
	PhoneNumber                string `json:"phone_number"`
	Email                      string `json:"email"`
	VerifyFlag                 bool   `json:"verify_flag"`
	RequiredChangePasswordFlag bool   `json:"required_change_password_flag"`
	CreatedAt                  string `json:"created_at"`
	UpdatedAt                  string `json:"updated_at"`
}

type TokenResponse struct {
	UserID                uuid.UUID `json:"user_id"`
	AccessToken           string    `json:"access_token"`
	AccessTokenExpiresAt  time.Time `json:"access_token_expires_at"`
	RefreshToken          string    `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time `json:"refresh_token_expires_at"`
}
