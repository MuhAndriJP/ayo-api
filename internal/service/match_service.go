package service

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

type MatchService interface {
	List(query dto.ListQuery) ([]*dto.MatchResponse, int64, error)
	GetByID(id int64) (*dto.MatchResponse, error)
	Create(req dto.CreateMatchRequest) error
	Update(id int64, req dto.UpdateMatchRequest) error
	Delete(id int64) error
	SaveResult(id int64, req dto.MatchResultRequest) error
}

type matchService struct {
	repo       repository.MatchRepository
	teamRepo   repository.TeamRepository
	playerRepo repository.PlayerRepository
}

func NewMatchService(repo repository.MatchRepository, teamRepo repository.TeamRepository, playerRepo repository.PlayerRepository) MatchService {
	return &matchService{repo: repo, teamRepo: teamRepo, playerRepo: playerRepo}
}

func (s *matchService) List(query dto.ListQuery) ([]*dto.MatchResponse, int64, error) {
	filter, err := query.ToFilter()
	if err != nil {
		return nil, 0, err
	}

	matches, total, err := s.repo.FindAll(filter)
	if err != nil {
		return nil, 0, err
	}

	return dto.MatchListToResponse(matches), total, nil
}

func (s *matchService) GetByID(id int64) (*dto.MatchResponse, error) {
	m, err := s.repo.FindByID(id)
	if err != nil {
		return nil, util.ErrNotFound(err, "pertandingan")
	}

	return dto.MatchToResponse(m), nil
}

func (s *matchService) Create(req dto.CreateMatchRequest) error {
	if req.HomeTeamID == req.AwayTeamID {
		return util.ErrSameTeam
	}

	if err := s.validateTeams(req.HomeTeamID, req.AwayTeamID); err != nil {
		return err
	}

	matchDate, err := util.ParseDate(req.MatchDate)
	if err != nil {
		return err
	}

	return s.repo.Create(req.ToModel(matchDate))
}

func (s *matchService) Update(id int64, req dto.UpdateMatchRequest) error {
	match, err := s.repo.FindByID(id)
	if err != nil {
		return util.ErrNotFound(err, "pertandingan")
	}

	if match.Status == model.MatchStatusFinished {
		return util.ErrMatchFinished
	}

	if req.HomeTeamID != 0 || req.AwayTeamID != 0 {
		if req.HomeTeamID == req.AwayTeamID {
			return util.ErrSameTeam
		}

		if err := s.validateTeams(req.HomeTeamID, req.AwayTeamID); err != nil {
			return err
		}

		match.HomeTeamID, match.AwayTeamID = req.HomeTeamID, req.AwayTeamID
	}

	if req.MatchDate != "" {
		t, err := util.ParseDate(req.MatchDate)
		if err != nil {
			return err
		}

		match.MatchDate = t
	}

	if req.MatchTime != "" {
		match.MatchTime = req.MatchTime
	}

	return s.repo.Update(match)
}

func (s *matchService) Delete(id int64) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return util.ErrNotFound(err, "pertandingan")
	}

	return s.repo.Delete(id)
}

func (s *matchService) SaveResult(id int64, req dto.MatchResultRequest) error {
	match, err := s.repo.FindByID(id)
	if err != nil {
		return util.ErrNotFound(err, "pertandingan")
	}

	goals, err := s.buildGoals(match, req.Goals)
	if err != nil {
		return err
	}

	match.HomeScore = req.HomeScore
	match.AwayScore = req.AwayScore
	match.Status = model.MatchStatusFinished

	return s.repo.SaveResult(match, goals)
}
