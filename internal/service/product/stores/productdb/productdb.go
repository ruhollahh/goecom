package productdb

import (
	"github.com/jmoiron/sqlx"
	"github.com/ruhollahh/go-ecom/pkg/logger"
)

type Store struct {
	log *logger.Logger
	db  sqlx.ExtContext
}

func NewStore(log *logger.Logger, db sqlx.ExtContext) *Store {
	return &Store{
		log: log,
		db:  db,
	}
}
