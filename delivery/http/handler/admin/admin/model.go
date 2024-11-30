package adminhandler

import (
	adminsvc "github.com/ruhollahh/go-ecom/internal/service/admin"
	"github.com/ruhollahh/go-ecom/pkg/validate"
)

type DlvLoginReq struct {
	PhoneNumber string `json:"phone_number" validate:"required"`
	Password    string `json:"password" validate:"required,gte=8,lte=64"`
}

type DlvLoginResp struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func toSvcLoginReq(req DlvLoginReq) adminsvc.LoginReq {
	return adminsvc.LoginReq{
		PhoneNumber: req.PhoneNumber,
		Password:    req.Password,
	}
}

func toDlvLoginResp(res adminsvc.LoginResp) DlvLoginResp {
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
