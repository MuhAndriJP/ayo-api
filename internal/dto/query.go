package dto

import "github.com/MuhAndriJP/ayo-api/internal/util"

type ListQuery struct {
	City    string `form:"city"`
	Search  string `form:"search"`
	Page    int64  `form:"page,default=1" binding:"min=1"`
	Limit   int64  `form:"limit,default=10" binding:"min=1,max=100"`
	SortBy  string `form:"sort_by"`
	SortDir string `form:"sort_dir"`
}

func (q ListQuery) ToFilter() (*ListFilter, error) {
	return &ListFilter{
		City:   q.City,
		Search: q.Search,
		Order:  util.BuildOrderClause(q.SortBy, q.SortDir),
		Offset: (q.Page - 1) * q.Limit,
		Limit:  q.Limit,
	}, nil
}

type ListFilter struct {
	City   string
	Search string
	Order  string
	Offset int64
	Limit  int64
}
