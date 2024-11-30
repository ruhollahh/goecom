package userhandler

import usersvc "github.com/ruhollahh/go-ecom/internal/service/user"

type Handler struct {
	userSvc *usersvc.Service
}

func New(userSvc *usersvc.Service) *Handler {
	return &Handler{
		userSvc: userSvc,
	}
}
