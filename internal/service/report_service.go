package service

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

type ReportService interface {
	GetMatchReport(matchID int64) (*dto.ReportResponse, error)
}

type reportService struct {
	matchRepo repository.MatchRepository
}

func NewReportService(matchRepo repository.MatchRepository) ReportService {
	return &reportService{matchRepo: matchRepo}
}

func (s *reportService) GetMatchReport(matchID int64) (*dto.ReportResponse, error) {
	match, err := s.matchRepo.FindByID(matchID)
	if err != nil {
		return nil, util.ErrNotFound(err, "pertandingan")
	}

	if match.Status != model.MatchStatusFinished {
		return nil, util.ErrMatchNotDone
	}

	homeWins, err := s.matchRepo.CountWins(match.HomeTeamID, match.MatchDate)
	if err != nil {
		return nil, err
	}

	awayWins, err := s.matchRepo.CountWins(match.AwayTeamID, match.MatchDate)
	if err != nil {
		return nil, err
	}

	return dto.ReportFromMatch(match, homeWins, awayWins), nil
}
