package userhandler

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	usersvc "github.com/ruhollahh/go-ecom/internal/service/user"
	"github.com/ruhollahh/go-ecom/pkg/expectederr"
	"github.com/ruhollahh/go-ecom/pkg/validate"
	"net/http"
)

// Signup godoc
// @Summary 	 User signup by PhoneNumber
// @Tags         Users
// @Accept       json
// @Produce      json
// @Param        Request body  DlvSignupReq true "User Signup request body"
// @Success      200  {object} DlvSignupResp
// @Failure      400  {object} expectederr.ErrorDocument{error=string}
// @Failure      422  {object} expectederr.ErrorDocument{error=string,fields=map[string]string}
// @Failure      500  {object} expectederr.ErrorDocument{error=string}
// @Router       /v1/users/signup [post].
func (h *Handler) Signup(c echo.Context) error {
	var req DlvSignupReq
	if err := c.Bind(&req); err != nil {
		return expectederr.NewError(err, http.StatusBadRequest)
	}
	if err := req.Validate(); err != nil {
		return expectederr.NewError(err, http.StatusUnprocessableEntity)
	}

	usr, err := h.userSvc.Signup(c.Request().Context(), toSvcSignupReq(req))
	if err != nil {
		if errors.Is(err, usersvc.ErrUniquePhoneNumber) {
			return expectederr.NewError(validate.NewFieldsError("phone_number", usersvc.ErrUniquePhoneNumber), http.StatusUnprocessableEntity)
		}

		return fmt.Errorf("signup: usr[%+v]: %w", usr, err)
	}

	return c.JSON(http.StatusOK, toDlvSignupResp(usr))
}
