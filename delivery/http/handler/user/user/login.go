package userhandler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	usersvc "github.com/ruhollahh/go-ecom/internal/service/user"
	"github.com/ruhollahh/go-ecom/pkg/expectederr"
	"net/http"
)

// Login godoc
// @Summary 	 User login by PhoneNumber
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        Request body  DlvLoginReq true "User login request body"
// @Success      200  {object} DlvLoginResp
// @Failure      400  {object} expectederr.ErrorDocument{error=string}
// @Failure      422  {object} expectederr.ErrorDocument{error=string,fields=map[string]string}
// @Failure      500  {object} expectederr.ErrorDocument{error=string}
// @Router       /v1/users/login [post].
func (h *Handler) Login(c echo.Context) error {
	var req DlvLoginReq
	if err := c.Bind(&req); err != nil {
		return expectederr.NewError(err, http.StatusBadRequest)
	}
	if err := req.Validate(); err != nil {
		return expectederr.NewError(err, http.StatusUnprocessableEntity)
	}

	usr, err := h.userSvc.Login(c.Request().Context(), toSvcLoginReq(req))
	if err != nil {
		if errors.Is(err, usersvc.ErrWrongCredentials) {
			return expectederr.NewError(usersvc.ErrWrongCredentials, http.StatusUnauthorized)
		}

		return fmt.Errorf("login: usr[%+v]: %w", usr, err)
	}

	return c.JSON(http.StatusCreated, toDlvLoginResp(usr))
}
