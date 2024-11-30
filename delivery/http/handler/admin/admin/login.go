package adminhandler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	adminsvc "github.com/ruhollahh/go-ecom/internal/service/admin"
	"github.com/ruhollahh/go-ecom/pkg/expectederr"
	"net/http"
)

// Login godoc
// @Summary 	 Admin login by PhoneNumber
// @Tags         Admins
// @Accept       json
// @Produce      json
// @Param        Request body  DlvLoginReq true "Admin login request body"
// @Success      200  {object} DlvLoginResp
// @Failure      400  {object} expectederr.ErrorDocument{error=string}
// @Failure      422  {object} expectederr.ErrorDocument{error=string,fields=map[string]string}
// @Failure      500  {object} expectederr.ErrorDocument{error=string}
// @Router       /v1/admins/login [post].
func (h *Handler) Login(c echo.Context) error {
	var req DlvLoginReq
	if err := c.Bind(&req); err != nil {
		return expectederr.NewError(err, http.StatusBadRequest)
	}
	if err := req.Validate(); err != nil {
		return expectederr.NewError(err, http.StatusUnprocessableEntity)
	}

	usr, err := h.adminSvc.Login(c.Request().Context(), toSvcLoginReq(req))
	if err != nil {
		if errors.Is(err, adminsvc.ErrWrongCredentials) {
			return expectederr.NewError(adminsvc.ErrWrongCredentials, http.StatusUnauthorized)
		}
		return fmt.Errorf("admin: adm[%+v]: %w", usr, err)
	}

	return c.JSON(http.StatusCreated, toDlvLoginResp(usr))
}
