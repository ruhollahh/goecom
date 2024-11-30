package admindb

import (
	"context"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"github.com/ruhollahh/go-ecom/internal/entity/admin"
	adminsvc "github.com/ruhollahh/go-ecom/internal/service/admin"
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

func (s *Store) QueryByPhoneNumber(ctx context.Context, phoneNumber string) (admin.Admin, error) {
	data := struct {
		PhoneNumber string `db:"phone_number"`
	}{
		PhoneNumber: phoneNumber,
	}

	const q = `
	SELECT
        id, name, phone_number, hashed_password, active, created_at, updated_at
	FROM
		admins
	WHERE
		phone_number = :phone_number`

	var dbUsr dbAdmin
	if err := dbpostgre.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbUsr); err != nil {
		if errors.Is(err, dbpostgre.ErrDBNotFound) {
			return admin.Admin{}, fmt.Errorf("namedquerystruct: %w", adminsvc.ErrNotFound)
		}
		return admin.Admin{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	usr, err := toCoreUser(dbUsr)
	if err != nil {
		return admin.Admin{}, err
	}

	return usr, nil
}
