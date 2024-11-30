package handler

import "github.com/ruhollahh/go-ecom/internal/entity/page"

type Result[T any] struct {
	Items       []T `json:"items"`
	Total       int `json:"total"`
	Page        int `json:"page"`
	RowsPerPage int `json:"rows_per_page"`
}

func NewResult[T any](items []T, total int, page page.Page) Result[T] {
	return Result[T]{
		Items:       items,
		Total:       total,
		Page:        page.Number(),
		RowsPerPage: page.RowsPerPage(),
	}
}
