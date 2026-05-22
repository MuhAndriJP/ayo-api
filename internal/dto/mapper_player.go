package dto

import (
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

func PlayerListToResponse(players []model.Player) []*PlayerResponse {
	res := make([]*PlayerResponse, len(players))
	for i := range players {
		res[i] = PlayerToResponse(&players[i])
	}
	return res
}

func PlayerToResponse(p *model.Player) *PlayerResponse {
	teamName := ""
	if p.Team.ID != 0 {
		teamName = p.Team.Name
	}
	return &PlayerResponse{
		ID:           int64(p.ID),
		TeamID:       p.TeamID,
		TeamName:     teamName,
		Name:         p.Name,
		HeightCm:     p.HeightCm,
		WeightKg:     p.WeightKg,
		Position:     p.Position.String(),
		JerseyNumber: p.JerseyNumber,
		CreatedAt:    util.FormatDateTime(p.CreatedAt),
	}
}
