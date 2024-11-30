package userdb

import (
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"github.com/ruhollahh/go-ecom/internal/entity/user"
	usersvc "github.com/ruhollahh/go-ecom/internal/service/user"
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

func (s *Store) Create(ctx context.Context, usr user.User) error {
	const q = `
	INSERT INTO users
		(id, name, phone_number, hashed_password, active, created_at, updated_at)
	VALUES
		(:id, :name, :phone_number, :hashed_password, :active, :created_at, :updated_at)`

	if err := dbpostgre.NamedExecContext(ctx, s.log, s.db, q, toDBUser(usr)); err != nil {
		if errors.Is(err, dbpostgre.ErrDBDuplicatedEntry) {
			return fmt.Errorf("namedexeccontext: %w", usersvc.ErrUniquePhoneNumber)
		}
		return fmt.Errorf("namedexeccontext: %w", err)
	}

	return nil
}

func (s *Store) QueryByID(ctx context.Context, userID uuid.UUID) (user.User, error) {
	data := struct {
		ID string `db:"id"`
	}{
		ID: userID.String(),
	}

	const q = `
	SELECT
        id, name, phone_number, hashed_password, active, created_at, updated_at
	FROM
		users
	WHERE 
		id = :id`

	var dbUsr dbUser
	if err := dbpostgre.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbUsr); err != nil {
		if errors.Is(err, dbpostgre.ErrDBNotFound) {
			return user.User{}, fmt.Errorf("namedquerystruct: %w", usersvc.ErrNotFound)
		}
		return user.User{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	usr, err := toCoreUser(dbUsr)
	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}

func (s *Store) QueryByPhoneNumber(ctx context.Context, phoneNumber string) (user.User, error) {
	data := struct {
		PhoneNumber string `db:"phone_number"`
	}{
		PhoneNumber: phoneNumber,
	}

	const q = `
	SELECT
        id, name, phone_number, hashed_password, active, created_at, updated_at
	FROM
		users
	WHERE
		phone_number = :phone_number`

	var dbUsr dbUser
	if err := dbpostgre.NamedQueryStruct(ctx, s.log, s.db, q, data, &dbUsr); err != nil {
		if errors.Is(err, dbpostgre.ErrDBNotFound) {
			return user.User{}, fmt.Errorf("namedquerystruct: %w", usersvc.ErrNotFound)
		}
		return user.User{}, fmt.Errorf("namedquerystruct: %w", err)
	}

	usr, err := toCoreUser(dbUsr)
	if err != nil {
		return user.User{}, err
	}

	return usr, nil
}
