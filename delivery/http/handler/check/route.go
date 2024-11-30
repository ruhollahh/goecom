package checkhandler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(e *echo.Group) {
	r := e.Group("/checks")

	r.GET("/readiness", h.Readiness)
	r.GET("/liveness", h.Liveness)
}
