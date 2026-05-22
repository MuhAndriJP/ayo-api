package service

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

func (s *matchService) buildGoals(match *model.Match, inputs []dto.GoalInput) ([]model.Goal, error) {
	goals := make([]model.Goal, 0, len(inputs))

	for _, g := range inputs {
		player, err := s.playerRepo.FindByID(g.PlayerID)
		if err != nil {
			return nil, util.ErrNotFound(err, "pemain")
		}

		if player.TeamID != match.HomeTeamID && player.TeamID != match.AwayTeamID {
			return nil, util.ErrPlayerNotInMatch
		}

		goals = append(goals, model.Goal{
			PlayerID:     g.PlayerID,
			MinuteScored: g.MinuteScored,
		})
	}

	return goals, nil
}

func (s *matchService) validateTeams(homeID, awayID int64) error {
	if _, err := s.teamRepo.FindByID(homeID); err != nil {
		return util.ErrNotFound(err, "tim home")
	}

	if _, err := s.teamRepo.FindByID(awayID); err != nil {
		return util.ErrNotFound(err, "tim away")
	}

	return nil
}
