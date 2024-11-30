package mid

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/delivery/http/claim"
	"github.com/ruhollahh/go-ecom/internal/service/auth"
)

func Auth(service authsvc.Service) echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) error {
			claims, err := service.Authenticate(ctx.Request().Header.Get("authorization"))
			if err != nil {
				return fmt.Errorf("authenticate: %w", err)
			}

			c := claim.SetClaims(ctx.Request().Context(), claims)
			ctx.SetRequest(ctx.Request().WithContext(c))

			return handler(ctx)
		}

		return h
	}

	return m
}
