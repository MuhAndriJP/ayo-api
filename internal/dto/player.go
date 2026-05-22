package dto

import "github.com/MuhAndriJP/ayo-api/internal/model"

func (r CreatePlayerRequest) ToModel(teamID int64) *model.Player {
	return &model.Player{TeamID: teamID, Name: r.Name, HeightCm: r.HeightCm, WeightKg: r.WeightKg, Position: model.PositionFromString(r.Position), JerseyNumber: r.JerseyNumber}
}

type CreatePlayerRequest struct {
	Name         string  `json:"name" binding:"required,min=2,max=200"`
	HeightCm     float32 `json:"height_cm" binding:"required,min=100,max=250"`
	WeightKg     float32 `json:"weight_kg" binding:"required,min=30,max=200"`
	Position     string  `json:"position" binding:"required,oneof=penyerang gelandang bertahan penjaga_gawang"`
	JerseyNumber int64   `json:"jersey_number" binding:"required,min=1,max=99"`
}

type UpdatePlayerRequest struct {
	Name         string  `json:"name" binding:"omitempty,min=2,max=200"`
	HeightCm     float32 `json:"height_cm" binding:"omitempty,min=100,max=250"`
	WeightKg     float32 `json:"weight_kg" binding:"omitempty,min=30,max=200"`
	Position     string  `json:"position" binding:"omitempty,oneof=penyerang gelandang bertahan penjaga_gawang"`
	JerseyNumber int64   `json:"jersey_number" binding:"omitempty,min=1,max=99"`
}


type PlayerResponse struct {
	ID           int64   `json:"id"`
	TeamID       int64   `json:"team_id"`
	TeamName     string  `json:"team_name,omitempty"`
	Name         string  `json:"name"`
	HeightCm     float32 `json:"height_cm"`
	WeightKg     float32 `json:"weight_kg"`
	Position     string  `json:"position"`
	JerseyNumber int64   `json:"jersey_number"`
	CreatedAt    string  `json:"created_at"`
}
