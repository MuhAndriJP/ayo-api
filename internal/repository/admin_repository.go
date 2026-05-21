package repository

import (
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"gorm.io/gorm"
)

type AdminRepository interface {
	FindByUsername(username string) (*model.Admin, error)
	Create(admin *model.Admin) error
}

type adminRepo struct{ db *gorm.DB }

func NewAdminRepository(db *gorm.DB) AdminRepository {
	return &adminRepo{db: db}
}

func (r *adminRepo) FindByUsername(username string) (*model.Admin, error) {
	var admin model.Admin
	err := r.db.Where("username = ?", username).First(&admin).Error
	return &admin, err
}

func (r *adminRepo) Create(admin *model.Admin) error {
	return r.db.Create(admin).Error
}
