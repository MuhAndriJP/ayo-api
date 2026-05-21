package dto

import (
	"fmt"
	"path/filepath"

	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

func TeamListToResponse(teams []model.Team) []*TeamResponse {
	res := make([]*TeamResponse, len(teams))
	for i := range teams {
		res[i] = TeamToResponse(&teams[i])
	}
	return res
}

func TeamToResponse(t *model.Team) *TeamResponse {
	logoURL := ""
	if t.LogoPath != "" {
		logoURL = fmt.Sprintf("/uploads/%s", filepath.Base(t.LogoPath))
	}
	return &TeamResponse{
		ID:          int64(t.ID),
		Name:        t.Name,
		LogoURL:     logoURL,
		FoundedYear: t.FoundedYear,
		Address:     t.Address,
		City:        t.City,
		CreatedAt:   util.FormatDateTime(t.CreatedAt),
	}
}
