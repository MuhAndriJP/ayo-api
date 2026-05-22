package dto

import (
	"time"

	"github.com/MuhAndriJP/ayo-api/internal/model"
	"github.com/MuhAndriJP/ayo-api/internal/util"
)

type ListQuery struct {
	City    string `form:"city"`
	Search  string `form:"search"`
	From    string `form:"from"`
	To      string `form:"to"`
	TeamID  int64  `form:"team_id"`
	Status  string `form:"status" binding:"omitempty,oneof=scheduled finished"`
	Page    int64  `form:"page,default=1" binding:"min=1"`
	Limit   int64  `form:"limit,default=10" binding:"min=1,max=100"`
	SortBy  string `form:"sort_by"`
	SortDir string `form:"sort_dir"`
}

func (q ListQuery) ToFilter() (*ListFilter, error) {
	from, err := util.ParseOptionalDate(q.From)
	if err != nil {
		return nil, err
	}

	to, err := util.ParseOptionalDate(q.To)
	if err != nil {
		return nil, err
	}

	return &ListFilter{
		City:   q.City,
		Search: q.Search,
		From:   from,
		To:     to,
		TeamID: q.TeamID,
		Status: model.MatchStatus(q.Status),
		Order:  util.BuildOrderClause(q.SortBy, q.SortDir),
		Offset: (q.Page - 1) * q.Limit,
		Limit:  q.Limit,
	}, nil
}

type ListFilter struct {
	City   string
	Search string
	From   time.Time
	To     time.Time
	TeamID int64
	Status model.MatchStatus
	Order  string
	Offset int64
	Limit  int64
}
