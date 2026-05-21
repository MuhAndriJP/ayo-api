package service

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/MuhAndriJP/ayo-api/internal/config"
	"github.com/MuhAndriJP/ayo-api/internal/dto"
	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/repository"
	"github.com/MuhAndriJP/ayo-api/internal/util"
	"github.com/google/uuid"
)

type TeamService interface {
	List(query dto.ListQuery) ([]*dto.TeamResponse, int64, error)
	GetByID(id int64) (*dto.TeamResponse, error)
	Create(req dto.CreateTeamRequest, logo *multipart.FileHeader) error
	Update(id int64, req dto.UpdateTeamRequest, logo *multipart.FileHeader) error
	Delete(id int64) error
}

type teamService struct {
	repo repository.TeamRepository
}

func NewTeamService(repo repository.TeamRepository) TeamService {
	return &teamService{repo: repo}
}

func (s *teamService) List(query dto.ListQuery) ([]*dto.TeamResponse, int64, error) {
	filter, err := query.ToFilter()
	if err != nil {
		return nil, 0, err
	}

	teams, total, err := s.repo.FindAll(filter)
	if err != nil {
		return nil, 0, err
	}

	return dto.TeamListToResponse(teams), total, nil
}

func (s *teamService) GetByID(id int64) (*dto.TeamResponse, error) {
	team, err := s.repo.FindByID(id)
	if err != nil {
		return nil, util.ErrNotFound(err, "tim")
	}

	return dto.TeamToResponse(team), nil
}

func (s *teamService) Create(req dto.CreateTeamRequest, logo *multipart.FileHeader) error {
	logoPath, err := saveLogoFile(logo)
	if err != nil {
		return err
	}

	return s.repo.Create(req.ToModel(logoPath))
}

func (s *teamService) Update(id int64, req dto.UpdateTeamRequest, logo *multipart.FileHeader) error {
	team, err := s.repo.FindByID(id)
	if err != nil {
		return util.ErrNotFound(err, "tim")
	}

	applyTeamFields(team, req)

	if logo != nil {
		path, err := saveLogoFile(logo)
		if err != nil {
			return err
		}
		team.LogoPath = path
	}

	return s.repo.Update(team)
}

func (s *teamService) Delete(id int64) error {
	if _, err := s.repo.FindByID(id); err != nil {
		return util.ErrNotFound(err, "tim")
	}

	return s.repo.Delete(id)
}

func applyTeamFields(team *model.Team, req dto.UpdateTeamRequest) {
	if req.Name != "" {
		team.Name = req.Name
	}

	if req.FoundedYear != 0 {
		team.FoundedYear = req.FoundedYear
	}

	if req.Address != "" {
		team.Address = req.Address
	}

	if req.City != "" {
		team.City = req.City
	}
}

func saveLogoFile(header *multipart.FileHeader) (string, error) {
	if header == nil {
		return "", nil
	}

	if header.Size > 2<<20 {
		return "", util.ErrLogoTooLarge
	}

	src, err := header.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	buf := make([]byte, 512)
	n, _ := src.Read(buf)
	if mime := detectMIME(buf[:n]); mime != "image/jpeg" && mime != "image/png" {
		return "", util.ErrLogoInvalidType
	}

	if _, err := src.Seek(0, io.SeekStart); err != nil {
		return "", err
	}

	dir := filepath.Join(config.App.UploadDir, "teams")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}

	dest := filepath.Join(dir, uuid.NewString()+filepath.Ext(header.Filename))
	out, err := os.Create(dest)
	if err != nil {
		return "", err
	}
	defer out.Close()

	_, err = io.Copy(out, src)
	return dest, err
}

func detectMIME(buf []byte) string {
	if len(buf) >= 2 && buf[0] == 0xFF && buf[1] == 0xD8 {
		return "image/jpeg"
	}

	if len(buf) >= 4 && buf[0] == 0x89 && buf[1] == 'P' && buf[2] == 'N' && buf[3] == 'G' {
		return "image/png"
	}

	return "application/octet-stream"
}
