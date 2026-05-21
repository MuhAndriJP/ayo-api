package dto

import "github.com/MuhAndriJP/ayo-api/internal/model"

func (r RegisterRequest) ToModel(passwordHash string) *model.Admin {
	return &model.Admin{Username: r.Username, PasswordHash: passwordHash}
}

type LoginRequest struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=6"`
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required,min=3,max=100"`
	Password string `json:"password" binding:"required,min=8"`
}

type AuthResponse struct {
	Token string `json:"token"`
}
