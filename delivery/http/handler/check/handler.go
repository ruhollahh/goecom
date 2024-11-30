package checkhandler

import (
	"github.com/jmoiron/sqlx"
)

type Handler struct {
	build string
	db    *sqlx.DB
}

func New(build string, db *sqlx.DB) *Handler {
	return &Handler{
		build: build,
		db:    db,
	}
}
