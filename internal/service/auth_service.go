package service

import (
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

type AuthService interface {
	Login(req dto.LoginRequest) (string, error)
	Register(req dto.RegisterRequest) error
}

type authService struct {
	adminRepo repository.AdminRepository
}

func NewAuthService(adminRepo repository.AdminRepository) AuthService {
	return &authService{adminRepo: adminRepo}
}

func (s *authService) Login(req dto.LoginRequest) (string, error) {
	admin, err := s.adminRepo.FindByUsername(req.Username)
	if err != nil {
		return "", util.ErrBadCredentials
	}

	if !util.CheckPassword(req.Password, admin.PasswordHash) {
		return "", util.ErrBadCredentials
	}

	return util.GenerateJWT(int64(admin.ID))
}

func (s *authService) Register(req dto.RegisterRequest) error {
	hash, err := util.HashPassword(req.Password)
	if err != nil {
		return err
	}

	return s.adminRepo.Create(req.ToModel(hash))
}
