package adminhandler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(e *echo.Group) {
	r := e.Group("/admins")

	r.POST("/login", h.Login)
}
