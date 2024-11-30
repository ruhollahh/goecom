package userhandler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) RegisterRoutes(e *echo.Group) {
	r := e.Group("/users")

	r.POST("/signup", h.Signup)
	r.POST("/login", h.Login)
}
