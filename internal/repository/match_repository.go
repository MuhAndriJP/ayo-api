package repository

import (
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"gorm.io/gorm"
)

type MatchRepository interface {
	FindAll(f *dto.ListFilter) ([]model.Match, int64, error)
	FindByID(id int64) (*model.Match, error)
	Create(match *model.Match) error
	Update(match *model.Match) error
	Delete(id int64) error
	SaveResult(match *model.Match, goals []model.Goal) error
	CountWins(teamID int64, beforeDate time.Time) (int64, error)
}

type matchRepo struct{ db *gorm.DB }

func NewMatchRepository(db *gorm.DB) MatchRepository {
	return &matchRepo{db: db}
}

func (r *matchRepo) FindAll(f *dto.ListFilter) ([]model.Match, int64, error) {
	var matches []model.Match
	var total int64

	q := r.db.Model(&model.Match{}).Preload("HomeTeam").Preload("AwayTeam")

	if !f.From.IsZero() {
		q = q.Where("match_date >= ?", f.From)
	}

	if !f.To.IsZero() {
		q = q.Where("match_date <= ?", f.To)
	}

	if f.TeamID > 0 {
		q = q.Where("home_team_id = ? OR away_team_id = ?", f.TeamID, f.TeamID)
	}

	if f.Status != "" {
		q = q.Where("status = ?", f.Status)
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := q.Offset(int(f.Offset)).Limit(int(f.Limit)).Order(f.Order).Find(&matches).Error
	return matches, total, err
}

func (r *matchRepo) FindByID(id int64) (*model.Match, error) {
	var match model.Match
	err := r.db.
		Preload("HomeTeam").
		Preload("AwayTeam").
		Preload("Goals").
		Preload("Goals.Player").
		First(&match, id).Error
	return &match, err
}

func (r *matchRepo) Create(match *model.Match) error {
	return r.db.Create(match).Error
}

func (r *matchRepo) Update(match *model.Match) error {
	return r.db.Save(match).Error
}

func (r *matchRepo) Delete(id int64) error {
	return r.db.Delete(&model.Match{}, id).Error
}

func (r *matchRepo) SaveResult(match *model.Match, goals []model.Goal) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("match_id = ?", match.ID).Delete(&model.Goal{}).Error; err != nil {
			return err
		}

		if err := tx.Save(match).Error; err != nil {
			return err
		}

		for i := range goals {
			goals[i].MatchID = int64(match.ID)
		}

		if len(goals) > 0 {
			if err := tx.Create(&goals).Error; err != nil {
				return err
			}
		}

		return nil
	})
}

// CountWins counts finished matches where teamID won, with match_date <= beforeDate
func (r *matchRepo) CountWins(teamID int64, beforeDate time.Time) (int64, error) {
	var count int64
	err := r.db.Model(&model.Match{}).
		Where("status = ? AND match_date <= ?", model.MatchStatusFinished, beforeDate).
		Where(
			r.db.Where("home_team_id = ? AND home_score > away_score", teamID).
				Or("away_team_id = ? AND away_score > home_score", teamID),
		).
		Count(&count).Error
	return count, err
}
