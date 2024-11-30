package mid

import (
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/pkg/web/meta"
	"time"
)

func Tracer() echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) error {
			v := httpmeta.Meta{
				TraceID: uuid.NewString(),
				Now:     time.Now().UTC(),
			}
			c := httpmeta.Set(ctx.Request().Context(), &v)
			ctx.SetRequest(ctx.Request().WithContext(c))

			return handler(ctx)
		}

		return h
	}

	return m
}
