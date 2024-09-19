package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique" validate:"required,min=3,max=32"`
	Password string `json:"password" validate:"required,min=8,max=32"`
	Role     string `json:"role" validate:"required"` // Add this line
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"newuser"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required" example:"newuser"`
	Password string `json:"password" binding:"required" example:"password123"`
}

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required" example:"refresh_token_value"`
}

type RefreshTokenResponse struct {
	AccessToken string `json:"access_token"`
}

type LogoutResponse struct {
	Message string `json:"message"`
}

const (
	RoleSuperAdmin = "super_admin"
	RoleAdmin      = "admin"
	RoleMod        = "mod"
	RoleUser       = "user"
)
