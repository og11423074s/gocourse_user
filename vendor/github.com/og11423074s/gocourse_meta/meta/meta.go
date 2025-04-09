package meta

import (
	"strconv"
)

type Meta struct {
	Page       int `json:"page"`
	PerPage    int `json:"per_page"`
	PageCount  int `json:"page_count"`
	TotalCount int `json:"total_count"`
}

func New(page, perPage, total int, pagLimDef string) (*Meta, error) {

	if perPage <= 0 {
		var err error
		perPage, err = strconv.Atoi(pagLimDef)
		if err != nil {
			return nil, err
		}
	}

	pageCount := 0
	if total >= 0 {
		pageCount = (total + perPage - 1) / perPage
		if page > pageCount {
			page = pageCount
		}
	}

	if page < 1 {
		page = 1
	}

	return &Meta{
		Page:       page,
		PerPage:    perPage,
		PageCount:  pageCount,
		TotalCount: total,
	}, nil
}

func (m *Meta) Offset() int {
	return (m.Page - 1) * m.PerPage
}

func (m *Meta) Limit() int {
	return m.PerPage
}
