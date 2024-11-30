package producthandler

import "github.com/labstack/echo/v4"

func (h *Handler) RegisterRoutes(e *echo.Group) {
	r := e.Group("/products")

	r.GET("", h.GetAll)
}
