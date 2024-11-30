package userdb

import (
	"github.com/google/uuid"
	"github.com/ruhollahh/go-ecom/internal/entity/user"
	"time"
)

type dbUser struct {
	ID             uuid.UUID `db:"id"`
	Name           string    `db:"name"`
	PhoneNumber    string    `db:"phone_number"`
	HashedPassword []byte    `db:"hashed_password"`
	Active         bool      `db:"active"`
	CreatedAt      time.Time `db:"created_at"`
	UpdatedAt      time.Time `db:"updated_at"`
}

func toDBUser(usr user.User) dbUser {
	return dbUser{
		ID:             usr.ID,
		Name:           usr.Name,
		PhoneNumber:    usr.PhoneNumber,
		HashedPassword: usr.HashedPassword,
		Active:         usr.Active,
		CreatedAt:      usr.CreatedAt.UTC(),
		UpdatedAt:      usr.UpdatedAt.UTC(),
	}
}

func toCoreUser(dbUsr dbUser) (user.User, error) {
	usr := user.User{
		ID:             dbUsr.ID,
		Name:           dbUsr.Name,
		PhoneNumber:    dbUsr.PhoneNumber,
		HashedPassword: dbUsr.HashedPassword,
		Active:         dbUsr.Active,
		CreatedAt:      dbUsr.CreatedAt.In(time.Local),
		UpdatedAt:      dbUsr.UpdatedAt.In(time.Local),
	}

	return usr, nil
}
