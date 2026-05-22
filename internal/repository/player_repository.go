package repository

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"gorm.io/gorm"
)

type PlayerRepository interface {
	FindByTeam(teamID int64, f *dto.ListFilter) ([]model.Player, error)
	FindByID(id int64) (*model.Player, error)
	JerseyExists(teamID int64, jersey int64) (bool, error)
	Create(player *model.Player) error
	Update(player *model.Player) error
	Delete(id int64) error
}

type playerRepo struct{ db *gorm.DB }

func NewPlayerRepository(db *gorm.DB) PlayerRepository {
	return &playerRepo{db: db}
}

func (r *playerRepo) FindByTeam(teamID int64, f *dto.ListFilter) ([]model.Player, error) {
	var players []model.Player
	err := r.db.Where("team_id = ?", teamID).Order(f.Order).Find(&players).Error
	return players, err
}

func (r *playerRepo) FindByID(id int64) (*model.Player, error) {
	var player model.Player
	err := r.db.Preload("Team").First(&player, id).Error
	return &player, err
}

func (r *playerRepo) JerseyExists(teamID int64, jersey int64) (bool, error) {
	var count int64
	q := r.db.Model(&model.Player{}).Where("team_id = ? AND jersey_number = ?", teamID, jersey)
	err := q.Count(&count).Error
	return count > 0, err
}

func (r *playerRepo) Create(player *model.Player) error {
	return r.db.Create(player).Error
}

func (r *playerRepo) Update(player *model.Player) error {
	return r.db.Save(player).Error
}

func (r *playerRepo) Delete(id int64) error {
	return r.db.Delete(&model.Player{}, id).Error
}
