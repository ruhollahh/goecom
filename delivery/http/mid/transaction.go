package mid

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/internal/entity/transaction"
	"github.com/ruhollahh/go-ecom/pkg/logger"
)

// ExecuteInTransation starts a transaction around all the storage calls within
// the scope of the handler function.
func ExecuteInTransation(log *logger.Logger, bgn transaction.Beginner) echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) error {
			hasCommited := false

			log.Info(ctx.Request().Context(), "BEGIN TRANSACTION")
			tx, err := bgn.Begin()
			if err != nil {
				return fmt.Errorf("BEGIN TRANSACTION: %w", err)
			}

			defer func() {
				if !hasCommited {
					log.Info(ctx.Request().Context(), "ROLLBACK TRANSACTION")
				}

				if err := tx.Rollback(); err != nil {
					if errors.Is(err, sql.ErrTxDone) {
						return
					}
					log.Info(ctx.Request().Context(), "ROLLBACK TRANSACTION", "ERROR", err)
				}
			}()

			c := transaction.Set(ctx.Request().Context(), tx)
			ctx.SetRequest(ctx.Request().WithContext(c))

			if err := handler(ctx); err != nil {
				return fmt.Errorf("EXECUTE TRANSACTION: %w", err)
			}

			log.Info(ctx.Request().Context(), "COMMIT TRANSACTION")
			if err := tx.Commit(); err != nil {
				return fmt.Errorf("COMMIT TRANSACTION: %w", err)
			}

			hasCommited = true

			return nil
		}

		return h
	}

	return m
}
