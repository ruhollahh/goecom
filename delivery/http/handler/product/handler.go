package producthandler

import productsvc "github.com/ruhollahh/go-ecom/internal/service/product"

type Handler struct {
	productSvc *productsvc.Service
}

func New(productSvc *productsvc.Service) *Handler {
	return &Handler{
		productSvc: productSvc,
	}
}
