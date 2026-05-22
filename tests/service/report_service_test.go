package service_test

import (
	"testing"
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// --- Mock MatchRepository ---

type mockMatchRepo struct {
	matches    map[int64]*model.Match
	winsByTeam map[int64]int64
}

func newMockMatchRepo() *mockMatchRepo {
	return &mockMatchRepo{
		matches:    map[int64]*model.Match{},
		winsByTeam: map[int64]int64{},
	}
}

func (m *mockMatchRepo) FindAll(f *dto.ListFilter) ([]model.Match, int64, error) {
	return nil, 0, nil
}
func (m *mockMatchRepo) FindByID(id int64) (*model.Match, error) {
	match, ok := m.matches[id]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return match, nil
}
func (m *mockMatchRepo) Create(match *model.Match) error { return nil }
func (m *mockMatchRepo) Update(match *model.Match) error { return nil }
func (m *mockMatchRepo) Delete(id int64) error           { return nil }
func (m *mockMatchRepo) SaveResult(match *model.Match, goals []model.Goal) error {
	return nil
}
func (m *mockMatchRepo) CountWins(teamID int64, beforeDate time.Time) (int64, error) {
	return m.winsByTeam[teamID], nil
}

var _ repository.MatchRepository = (*mockMatchRepo)(nil)

func makeFinishedMatch(homeScore, awayScore int64) *model.Match {
	h := homeScore
	a := awayScore
	return &model.Match{
		Model:      gorm.Model{ID: 1},
		MatchDate:  time.Date(2026, 5, 10, 0, 0, 0, 0, time.UTC),
		MatchTime:  "15:00",
		HomeTeamID: 1,
		HomeTeam:   model.Team{Model: gorm.Model{ID: 1}, Name: "Tim A"},
		AwayTeamID: 2,
		AwayTeam:   model.Team{Model: gorm.Model{ID: 2}, Name: "Tim B"},
		HomeScore:  h,
		AwayScore:  a,
		Status:     model.MatchStatusFinished,
		Goals: []model.Goal{
			{Model: gorm.Model{ID: 1}, PlayerID: 10, Player: model.Player{Name: "Budi"}, MinuteScored: 15},
			{Model: gorm.Model{ID: 2}, PlayerID: 10, Player: model.Player{Name: "Budi"}, MinuteScored: 45},
			{Model: gorm.Model{ID: 3}, PlayerID: 11, Player: model.Player{Name: "Andi"}, MinuteScored: 72},
		},
	}
}

func TestReport_HomeWin(t *testing.T) {
	repo := newMockMatchRepo()
	repo.matches[1] = makeFinishedMatch(2, 1)
	repo.winsByTeam[1] = 5
	repo.winsByTeam[2] = 3

	svc := service.NewReportService(repo)
	report, err := svc.GetMatchReport(1)

	assert.NoError(t, err)
	assert.Equal(t, "Tim Home Menang", report.ResultStatus)
	assert.Equal(t, int64(2), report.FinalScore.Home)
	assert.Equal(t, int64(1), report.FinalScore.Away)
	assert.Equal(t, int64(5), report.HomeTeamTotalWinsUntilThisMatch)
	assert.Equal(t, int64(3), report.AwayTeamTotalWinsUntilThisMatch)
}

func TestReport_TopScorer(t *testing.T) {
	repo := newMockMatchRepo()
	repo.matches[1] = makeFinishedMatch(3, 0)

	svc := service.NewReportService(repo)
	report, err := svc.GetMatchReport(1)

	assert.NoError(t, err)
	assert.NotNil(t, report.TopScorer)
	assert.Equal(t, int64(10), report.TopScorer.PlayerID)
	assert.Equal(t, int64(2), report.TopScorer.Goals)
	assert.Equal(t, "Budi", report.TopScorer.Name)
}

func TestReport_Draw(t *testing.T) {
	repo := newMockMatchRepo()
	repo.matches[1] = makeFinishedMatch(1, 1)

	svc := service.NewReportService(repo)
	report, err := svc.GetMatchReport(1)

	assert.NoError(t, err)
	assert.Equal(t, "Draw", report.ResultStatus)
}

func TestReport_AwayWin(t *testing.T) {
	repo := newMockMatchRepo()
	repo.matches[1] = makeFinishedMatch(0, 2)

	svc := service.NewReportService(repo)
	report, err := svc.GetMatchReport(1)

	assert.NoError(t, err)
	assert.Equal(t, "Tim Away Menang", report.ResultStatus)
}

func TestReport_MatchNotFound(t *testing.T) {
	repo := newMockMatchRepo()
	svc := service.NewReportService(repo)
	_, err := svc.GetMatchReport(999)
	assert.Error(t, err)
	assert.Equal(t, "pertandingan tidak ditemukan", err.Error())
}

func TestReport_MatchNotFinished(t *testing.T) {
	repo := newMockMatchRepo()
	repo.matches[1] = &model.Match{
		Model:  gorm.Model{ID: 1},
		Status: model.MatchStatusScheduled,
	}
	svc := service.NewReportService(repo)
	_, err := svc.GetMatchReport(1)
	assert.Error(t, err)
	assert.Equal(t, "pertandingan belum selesai", err.Error())
}
