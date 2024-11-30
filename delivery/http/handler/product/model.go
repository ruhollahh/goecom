package producthandler

import (
	"github.com/google/uuid"
	"github.com/ruhollahh/go-ecom/internal/entity/product"
	"time"
)

type QueryParams struct {
	Page           string `query:"page"`
	Rows           string `query:"rows"`
	SortField      string `query:"sort_field"`
	SortDirection  string `query:"sort_direction"`
	ID             string `query:"id"`
	Name           string `query:"name"`
	Price          string `query:"price"`
	StartCreatedAt string `query:"start_created_at"`
	EndCreatedAt   string `query:"end_created_at"`
}

type DlvProduct struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageID     uuid.UUID `json:"image_id"`
	Price       int       `json:"price"`
	CreatedAt   time.Time `json:"created_at"`
}

func toDlvProduct(prd product.Product) DlvProduct {
	return DlvProduct{
		ID:          prd.ID,
		Name:        prd.Name,
		Description: prd.Description,
		ImageID:     prd.ImageID,
		Price:       prd.Price,
		CreatedAt:   prd.CreatedAt,
	}
}

func toDlvProducts(prds []product.Product) []DlvProduct {
	dlv := make([]DlvProduct, len(prds))
	for i, prd := range prds {
		dlv[i] = toDlvProduct(prd)
	}

	return dlv
}
