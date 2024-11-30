package producthandler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/ruhollahh/go-ecom/delivery/http/handler"
	"github.com/ruhollahh/go-ecom/pkg/expectederr"
	"net/http"
)

// GetAll godoc
// @Summary 	 Get all products
// @Tags         Products
// @Accept       json
// @Produce      json
// @Param        id query string false "Filter by id"
// @Param        name query string false "Filter by name"
// @Param        price query string false "Filter by price"
// @Param        start_created_at query string false "Filter by start created at"
// @Param        end_created_at query string false "Filter by end created at"
// @Param        page query string false "Page number"
// @Param        rows query string false "Page size"
// @Param        sort_field query string false "Sort by field" Enums(id,name,price)
// @Param        sort_direction query string false "Sort order" Enums(asc,desc)
// @Success      200  {object} QueryParams
// @Failure      400  {object} expectederr.ErrorDocument{error=string}
// @Failure      422  {object} expectederr.ErrorDocument{error=string,fields=map[string]string}
// @Failure      500  {object} expectederr.ErrorDocument{error=string}
// @Router       /v1/products [get].
func (h *Handler) GetAll(c echo.Context) error {
	var qp QueryParams
	if err := c.Bind(&qp); err != nil {
		return expectederr.NewError(err, http.StatusBadRequest)
	}

	pagination, err := parsePage(qp)
	if err != nil {
		return expectederr.NewError(err, http.StatusUnprocessableEntity)
	}

	filter, err := parseFilter(qp)
	if err != nil {
		return expectederr.NewError(err, http.StatusUnprocessableEntity)
	}

	sortBy, err := parseSort(qp)
	if err != nil {
		return expectederr.NewError(err, http.StatusUnprocessableEntity)
	}

	prds, err := h.productSvc.Query(c.Request().Context(), filter, sortBy, pagination)
	if err != nil {
		return fmt.Errorf("query: %w", err)
	}

	total, err := h.productSvc.Count(c.Request().Context(), filter)
	if err != nil {
		return fmt.Errorf("count: %w", err)
	}

	return c.JSON(http.StatusOK, handler.NewResult(toDlvProducts(prds), total, pagination))
}
