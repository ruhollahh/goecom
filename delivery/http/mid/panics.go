package mid

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/pkg/web/metrics"
	"runtime/debug"
)

// Panics recovers from panics and converts the panic to an error so it is
// reported in Metrics and handled in Errors.
func Panics() echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) (err error) {

			// Defer a function to recover from a panic and set the err return
			// variable after the fact.
			defer func() {
				if rec := recover(); rec != nil {
					trace := debug.Stack()
					err = fmt.Errorf("PANIC [%v] TRACE[%s]", rec, string(trace))

					metrics.AddPanics(ctx.Request().Context())
				}
			}()

			return handler(ctx)
		}

		return h
	}

	return m
}
