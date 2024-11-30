package commands

import (
	"context"
	"errors"
	"fmt"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre/migrate"
	"time"
)

// ErrHelp provides context that help was given.
var ErrHelp = errors.New("provided help")

// Migrate creates the schema in the database.
func Migrate(cfg dbpostgre.Config) error {
	db, err := dbpostgre.Open(cfg)
	if err != nil {
		return fmt.Errorf("connect database: %w", err)
	}
	defer db.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := migrate.Migrate(ctx, db); err != nil {
		return fmt.Errorf("migrate database: %w", err)
	}

	fmt.Println("migrations complete")
	return nil
}
