package adminhandler

import adminsvc "github.com/ruhollahh/go-ecom/internal/service/admin"

type Handler struct {
	adminSvc *adminsvc.Service
}

func New(adminService *adminsvc.Service) *Handler {
	return &Handler{
		adminSvc: adminService,
	}
}
