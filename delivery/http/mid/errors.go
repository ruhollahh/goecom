package mid

import (
	"errors"
	"github.com/labstack/echo/v4"
	authsvc "github.com/ruhollahh/go-ecom/internal/service/auth"
	"github.com/ruhollahh/go-ecom/pkg/expectederr"
	"github.com/ruhollahh/go-ecom/pkg/logger"
	"github.com/ruhollahh/go-ecom/pkg/validate"
	"net/http"
)

// Errors handles errors coming out of the call chain. It detects normal
// application errors which are used to respond to the client in a uniform way.
// Unexpected errors (status >= 500) are logged.
func Errors(log *logger.Logger) echo.MiddlewareFunc {
	m := func(handler echo.HandlerFunc) echo.HandlerFunc {
		h := func(ctx echo.Context) error {
			if err := handler(ctx); err != nil {
				var er expectederr.ErrorDocument
				var status int

				switch {
				case expectederr.IsError(err):
					reqErr := expectederr.GetError(err)
					if validate.IsFieldErrors(reqErr.Err) {
						fieldErrors := validate.GetFieldErrors(reqErr.Err)
						er = expectederr.ErrorDocument{
							Error:  "data validation error",
							Fields: fieldErrors.Fields(),
						}
						status = reqErr.Status
						break
					}
					er = expectederr.ErrorDocument{
						Error: reqErr.Error(),
					}
					status = reqErr.Status
				case errors.Is(err, authsvc.ErrAuthFailed):
					er = expectederr.ErrorDocument{
						Error: http.StatusText(http.StatusUnauthorized),
					}
					status = http.StatusUnauthorized
				case errors.Is(err, echo.ErrNotFound):
					er = expectederr.ErrorDocument{
						Error: http.StatusText(http.StatusNotFound),
					}
					status = http.StatusNotFound

				default:
					log.Error(ctx.Request().Context(), "message", "msg", err)
					er = expectederr.ErrorDocument{
						Error: http.StatusText(http.StatusInternalServerError),
					}
					status = http.StatusInternalServerError
				}

				if err = ctx.JSON(status, er); err != nil {
					return err
				}
			}

			return nil
		}

		return h
	}

	return m
}
