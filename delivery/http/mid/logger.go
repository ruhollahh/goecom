package mid

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/pkg/logger"
	"github.com/ruhollahh/go-ecom/pkg/web/meta"
	"time"
)

func Logger(log *logger.Logger) echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) error {
			v := httpmeta.Get(ctx.Request().Context())

			path := ctx.Request().URL.Path
			if ctx.Request().URL.RawQuery != "" {
				path = fmt.Sprintf("%s?%s", path, ctx.Request().URL.RawQuery)
			}

			log.Info(ctx.Request().Context(), "request started", "method", ctx.Request().Method, "path", path,
				"remoteaddr", ctx.Request().RemoteAddr)

			err := handler(ctx)

			log.Info(ctx.Request().Context(), "request completed", "method", ctx.Request().Method, "path", path,
				"remoteaddr", ctx.Request().RemoteAddr, "statuscode", v.StatusCode, "since", time.Since(v.Now))

			return err
		}

		return h
	}

	return m
}
