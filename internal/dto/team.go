package dto

import "github.com/MuhAndriJP/ayo-api/internal/model"

func (r CreateTeamRequest) ToModel(logoPath string) *model.Team {
	return &model.Team{Name: r.Name, LogoPath: logoPath, FoundedYear: r.FoundedYear, Address: r.Address, City: r.City}
}

type CreateTeamRequest struct {
	Name        string `form:"name" binding:"required,min=2,max=200"`
	FoundedYear int64  `form:"founded_year" binding:"required,min=1800,max=2026"`
	Address     string `form:"address" binding:"required"`
	City        string `form:"city" binding:"required,min=2,max=100"`
}

type UpdateTeamRequest struct {
	Name        string `form:"name" binding:"omitempty,min=2,max=200"`
	FoundedYear int64  `form:"founded_year" binding:"omitempty,min=1800,max=2026"`
	Address     string `form:"address" binding:"omitempty"`
	City        string `form:"city" binding:"omitempty,min=2,max=100"`
}

type TeamResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	LogoURL     string `json:"logo_url"`
	FoundedYear int64  `json:"founded_year"`
	Address     string `json:"address"`
	City        string `json:"city"`
	CreatedAt   string `json:"created_at"`
}

