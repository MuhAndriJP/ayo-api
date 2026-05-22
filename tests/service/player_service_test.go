package service_test

import (
	"errors"
	"testing"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/service"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

// --- Mock PlayerRepository ---

type mockPlayerRepo struct {
	players      map[int64]*model.Player
	jerseyExists bool
}

func newMockPlayerRepo() *mockPlayerRepo {
	return &mockPlayerRepo{players: map[int64]*model.Player{}}
}

func (m *mockPlayerRepo) FindByTeam(teamID int64, f *dto.ListFilter) ([]model.Player, error) {
	var res []model.Player
	for _, p := range m.players {
		if p.TeamID == teamID {
			res = append(res, *p)
		}
	}
	return res, nil
}

func (m *mockPlayerRepo) FindByID(id int64) (*model.Player, error) {
	p, ok := m.players[id]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return p, nil
}

func (m *mockPlayerRepo) JerseyExists(teamID int64, jersey int64) (bool, error) {
	return m.jerseyExists, nil
}

func (m *mockPlayerRepo) Create(player *model.Player) error {
	player.Model.ID = uint(len(m.players) + 1)
	m.players[int64(player.ID)] = player
	return nil
}

func (m *mockPlayerRepo) Update(player *model.Player) error {
	m.players[int64(player.ID)] = player
	return nil
}

func (m *mockPlayerRepo) Delete(id int64) error {
	delete(m.players, id)
	return nil
}

// --- Mock TeamRepository ---

type mockTeamRepo struct {
	teams map[int64]*model.Team
}

func newMockTeamRepo() *mockTeamRepo {
	return &mockTeamRepo{teams: map[int64]*model.Team{
		1: {Model: gorm.Model{ID: 1}, Name: "Tim A", City: "Jakarta"},
	}}
}

func (m *mockTeamRepo) FindAll(f *dto.ListFilter) ([]model.Team, int64, error) {
	return nil, 0, nil
}
func (m *mockTeamRepo) FindByID(id int64) (*model.Team, error) {
	t, ok := m.teams[id]
	if !ok {
		return nil, gorm.ErrRecordNotFound
	}
	return t, nil
}
func (m *mockTeamRepo) Create(team *model.Team) error { return nil }
func (m *mockTeamRepo) Update(team *model.Team) error { return nil }
func (m *mockTeamRepo) Delete(id int64) error         { return nil }

// --- Tests ---

func TestCreatePlayer_Success(t *testing.T) {
	playerRepo := newMockPlayerRepo()
	teamRepo := newMockTeamRepo()
	svc := service.NewPlayerService(playerRepo, teamRepo)

	req := dto.CreatePlayerRequest{
		Name:         "Budi Santoso",
		HeightCm:     175,
		WeightKg:     70,
		Position:     "penyerang",
		JerseyNumber: 9,
	}
	err := svc.Create(1, req)
	assert.NoError(t, err)
}

func TestCreatePlayer_TeamNotFound(t *testing.T) {
	playerRepo := newMockPlayerRepo()
	teamRepo := newMockTeamRepo()
	svc := service.NewPlayerService(playerRepo, teamRepo)

	req := dto.CreatePlayerRequest{
		Name:         "Test",
		HeightCm:     170,
		WeightKg:     65,
		Position:     "gelandang",
		JerseyNumber: 10,
	}
	err := svc.Create(99, req)
	assert.Error(t, err)
	assert.Equal(t, "tim tidak ditemukan", err.Error())
}

func TestCreatePlayer_DuplicateJersey(t *testing.T) {
	playerRepo := newMockPlayerRepo()
	playerRepo.jerseyExists = true
	teamRepo := newMockTeamRepo()
	svc := service.NewPlayerService(playerRepo, teamRepo)

	req := dto.CreatePlayerRequest{
		Name:         "Andi",
		HeightCm:     168,
		WeightKg:     62,
		Position:     "bertahan",
		JerseyNumber: 5,
	}
	err := svc.Create(1, req)
	assert.Error(t, err)
	assert.Equal(t, "nomor punggung sudah dipakai pemain lain di tim ini", err.Error())
}

func TestGetPlayer_NotFound(t *testing.T) {
	playerRepo := newMockPlayerRepo()
	teamRepo := newMockTeamRepo()
	svc := service.NewPlayerService(playerRepo, teamRepo)

	_, err := svc.GetByID(999)
	assert.Error(t, err)
	assert.Equal(t, "pemain tidak ditemukan", err.Error())
}

func TestDeletePlayer_NotFound(t *testing.T) {
	playerRepo := newMockPlayerRepo()
	teamRepo := newMockTeamRepo()
	svc := service.NewPlayerService(playerRepo, teamRepo)

	err := svc.Delete(999)
	assert.Error(t, err)
}

func TestUpdatePlayer_JerseyConflict(t *testing.T) {
	playerRepo := newMockPlayerRepo()
	playerRepo.players[1] = &model.Player{
		Model:        gorm.Model{ID: 1},
		TeamID:       1,
		Name:         "Saya",
		JerseyNumber: 7,
		Position:     model.PositionGelandang,
		HeightCm:     170,
		WeightKg:     65,
	}
	playerRepo.jerseyExists = true
	teamRepo := newMockTeamRepo()
	svc := service.NewPlayerService(playerRepo, teamRepo)

	req := dto.UpdatePlayerRequest{JerseyNumber: 10}
	err := svc.Update(1, req)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "nomor punggung")
}

// Keep errors import used
var _ = errors.New
