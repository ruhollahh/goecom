package product

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID          uuid.UUID
	Name        string
	Description string
	ImageID     uuid.UUID
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
