package productsvc

import (
	"github.com/google/uuid"
	"time"
)

type Filter struct {
	ID             *uuid.UUID
	Name           *string
	Price          *int
	StartCreatedAt *time.Time
	EndCreatedAt   *time.Time
}
