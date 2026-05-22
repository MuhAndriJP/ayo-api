package service

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

type PlayerService interface {
	ListByTeam(teamID int64, query dto.ListQuery) ([]*dto.PlayerResponse, error)
	GetByID(id int64) (*dto.PlayerResponse, error)
	Create(teamID int64, req dto.CreatePlayerRequest) error
	Update(id int64, req dto.UpdatePlayerRequest) error
	Delete(id int64) error
}

type playerService struct {
	repo     repository.PlayerRepository
	teamRepo repository.TeamRepository
}

func NewPlayerService(repo repository.PlayerRepository, teamRepo repository.TeamRepository) PlayerService {
	return &playerService{repo: repo, teamRepo: teamRepo}
}

func (s *playerService) ListByTeam(teamID int64, query dto.ListQuery) ([]*dto.PlayerResponse, error) {
	if _, err := s.teamRepo.FindByID(teamID); err != nil {
		return nil, util.ErrNotFound(err, "tim")
	}

	filter, err := query.ToFilter()
	if err != nil {
		return nil, err
	}

	players, err := s.repo.FindByTeam(teamID, filter)
	if err != nil {
		return nil, err
	}

	return dto.PlayerListToResponse(players), nil
}

func (s *playerService) GetByID(id int64) (*dto.PlayerResponse, error) {
	p, err := s.repo.FindByID(id)
	if err != nil {
		return nil, util.ErrNotFound(err, "pemain")
	}

	return dto.PlayerToResponse(p), nil
}

func (s *playerService) Create(teamID int64, req dto.CreatePlayerRequest) error {
	if _, err := s.teamRepo.FindByID(teamID); err != nil {
		return util.ErrNotFound(err, "tim")
	}

	isExists, err := s.repo.JerseyExists(teamID, req.JerseyNumber)
	if err != nil {
		return err
	}

	if isExists {
		return util.ErrJerseyExists
	}

	return s.repo.Create(req.ToModel(teamID))
}

func (s *playerService) Update(id int64, req dto.UpdatePlayerRequest) error {
	player, err := s.repo.FindByID(id)
	if err != nil {
		return util.ErrNotFound(err, "pemain")
	}

	if req.JerseyNumber != 0 && req.JerseyNumber != player.JerseyNumber {
		isExists, err := s.repo.JerseyExists(player.TeamID, req.JerseyNumber)
		if err != nil {
			return err
		}

		if isExists {
			return util.ErrJerseyExists
		}

		player.JerseyNumber = req.JerseyNumber
	}

	if req.Name != "" {
		player.Name = req.Name
	}

	if req.HeightCm != 0 {
		player.HeightCm = req.HeightCm
	}

	if req.WeightKg != 0 {
		player.WeightKg = req.WeightKg
	}

	if req.Position != "" {
		player.Position = model.PositionFromString(req.Position)
	}

	return s.repo.Update(player)
}

func (s *playerService) Delete(id int64) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return util.ErrNotFound(err, "pemain")
	}

	return s.repo.Delete(id)
}
