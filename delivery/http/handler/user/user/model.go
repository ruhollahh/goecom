package userhandler

import (
	"github.com/google/uuid"
	usersvc "github.com/ruhollahh/go-ecom/internal/service/user"
	"github.com/ruhollahh/go-ecom/pkg/validate"
)

type DlvSignupReq struct {
	Name            string `json:"name" validate:"required,lte=48"`
	PhoneNumber     string `json:"phone_number" validate:"required"`
	Password        string `json:"password" validate:"required,gte=8,lte=64"`
	PasswordConfirm string `json:"password_confirm" validate:"eqfield=Password"`
}

type DlvSignupResp struct {
	UserID uuid.UUID `json:"user_id"`
}

func toSvcSignupReq(req DlvSignupReq) usersvc.SignupReq {
	return usersvc.SignupReq{
		Name:        req.Name,
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
}

func toDlvSignupResp(res usersvc.SignupResp) DlvSignupResp {
	return DlvSignupResp{
		UserID: res.ID,
	}
}

func (r DlvSignupReq) Validate() error {
	if err := validate.Check(r); err != nil {
		return err
	}

	return nil
}

// =============================================================================

type DlvLoginReq struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required,gte=8,lte=64"`
}

type DlvLoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func toSvcLoginReq(req DlvLoginReq) usersvc.LoginReq {
	return usersvc.LoginReq{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
}

func toDlvLoginResp(res usersvc.LoginResp) DlvLoginResp {
	return DlvLoginResp{
		AccessToken:  res.AccessToken,
		RefreshToken: res.RefreshToken,
	}
}

func (r DlvLoginReq) Validate() error {
	if err := validate.Check(r); err != nil {
		return err
	}

	return nil
}
