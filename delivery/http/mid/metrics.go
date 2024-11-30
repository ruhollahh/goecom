package mid

import (
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/pkg/web/metrics"
)

// Metrics updates program counters.
func Metrics() echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) error {
			c := metrics.Set(ctx.Request().Context())
			ctx.SetRequest(ctx.Request().WithContext(c))

			n := metrics.AddRequests(ctx.Request().Context())
			if n%1000 == 0 {
				metrics.AddGoroutines(ctx.Request().Context())
			}

			err := handler(ctx)
			if err != nil {
				metrics.AddErrors(ctx.Request().Context())
			}

			return err
		}

		return h
	}

	return m
}
