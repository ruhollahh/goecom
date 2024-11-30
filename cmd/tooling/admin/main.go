// This program performs administrative tasks for the garage sale service.
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/ardanlabs/conf/v3"
	"github.com/ruhollahh/go-ecom/cmd/tooling/admin/commands"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"github.com/ruhollahh/go-ecom/pkg/logger"
	"io"
	"os"
	"time"
)

var build = "develop"

type config struct {
	conf.Version
	Args     conf.Args
	Postgres struct {
		User         string
		Password     string
		Host         string
		Port         string
		Name         string
		MaxIdleConns int           `conf:"default:25"`
		MaxOpenConns int           `conf:"default:25"`
		MaxIdleTime  time.Duration `conf:"default:15m"`
		DisableTLS   bool          `conf:"default:true"`
	}
}

func main() {
	log := logger.New(io.Discard, logger.LevelInfo, "ADMIN", func(context.Context) string { return "00000000-0000-0000-0000-000000000000" })

	if err := run(log); err != nil {
		if !errors.Is(err, commands.ErrHelp) {
			fmt.Println("msg", err)
		}
		os.Exit(1)
	}
}

func run(log *logger.Logger) error {
	cfg := config{
		Version: conf.Version{
			Build: build,
			Desc:  "copyright information here",
		},
	}

	const prefix = "GOECOM"
	help, err := conf.Parse(prefix, &cfg)
	if err != nil {
		if errors.Is(err, conf.ErrHelpWanted) {
			fmt.Println(help)
			return nil
		}

		out, err := conf.String(&cfg)
		if err != nil {
			return fmt.Errorf("generating config for output: %w", err)
		}
		log.Info(context.Background(), "startup", "config", out)

		return fmt.Errorf("parsing config: %w", err)
	}

	return processCommands(cfg.Args, cfg)
}

// processCommands handles the execution of the commands specified on
// the command line.
func processCommands(args conf.Args, cfg config) error {
	dbConfig := dbpostgre.Config{
		User:         cfg.Postgres.User,
		Password:     cfg.Postgres.Password,
		Host:         cfg.Postgres.Host,
		Port:         cfg.Postgres.Port,
		Name:         cfg.Postgres.Name,
		MaxIdleConns: cfg.Postgres.MaxIdleConns,
		MaxOpenConns: cfg.Postgres.MaxOpenConns,
		MaxIdleTime:  cfg.Postgres.MaxIdleTime,
	}

	switch args.Num(0) {
	case "migrate":
		if err := commands.Migrate(dbConfig); err != nil {
			return fmt.Errorf("migrating database: %w", err)
		}

	case "seed":
		if err := commands.Seed(dbConfig); err != nil {
			return fmt.Errorf("seeding database: %w", err)
		}

	case "migrate-seed":
		if err := commands.Migrate(dbConfig); err != nil {
			return fmt.Errorf("migrating database: %w", err)
		}
		if err := commands.Seed(dbConfig); err != nil {
			return fmt.Errorf("seeding database: %w", err)
		}

	default:
		fmt.Println("migrate:    create the schema in the database")
		fmt.Println("seed:       add data to the database")
		fmt.Println("provide a command to get more help.")
		return commands.ErrHelp
	}

	return nil
}
