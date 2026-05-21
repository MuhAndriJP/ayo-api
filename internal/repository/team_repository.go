package repository

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"gorm.io/gorm"
)

type TeamRepository interface {
	FindAll(f *dto.ListFilter) ([]model.Team, int64, error)
	FindByID(id int64) (*model.Team, error)
	Create(team *model.Team) error
	Update(team *model.Team) error
	Delete(id int64) error
}

type teamRepo struct{ db *gorm.DB }

func NewTeamRepository(db *gorm.DB) TeamRepository {
	return &teamRepo{db: db}
}

func (r *teamRepo) FindAll(f *dto.ListFilter) ([]model.Team, int64, error) {
	var teams []model.Team
	var total int64

	q := r.db.Model(&model.Team{})
	if f.City != "" {
		q = q.Where("city = ?", f.City)
	}

	if f.Search != "" {
		q = q.Where("name LIKE ?", "%"+f.Search+"%")
	}

	if err := q.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	err := q.Offset(int(f.Offset)).Limit(int(f.Limit)).Order(f.Order).Find(&teams).Error
	return teams, total, err
}

func (r *teamRepo) FindByID(id int64) (*model.Team, error) {
	var team model.Team
	err := r.db.First(&team, id).Error
	return &team, err
}

func (r *teamRepo) Create(team *model.Team) error {
	return r.db.Create(team).Error
}

func (r *teamRepo) Update(team *model.Team) error {
	return r.db.Save(team).Error
}

func (r *teamRepo) Delete(id int64) error {
	return r.db.Delete(&model.Team{}, id).Error
}
