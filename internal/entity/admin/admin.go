package admin

import (
	"github.com/google/uuid"
	"time"
)

type Admin struct {
	ID             uuid.UUID
	Name           string
	PhoneNumber    string
	HashedPassword []byte
	Active         bool
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
