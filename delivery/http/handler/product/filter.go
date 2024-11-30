package producthandler

import (
	"github.com/google/uuid"
	productsvc "github.com/ruhollahh/go-ecom/internal/service/product"
	"github.com/ruhollahh/go-ecom/pkg/validate"
	"strconv"
	"time"
)

func parseFilter(qp QueryParams) (productsvc.Filter, error) {
	var filter productsvc.Filter

	if qp.ID != "" {
		id, err := uuid.Parse(qp.ID)
		if err != nil {
			return productsvc.Filter{}, validate.NewFieldsError("id", err)
		}
		filter.ID = &id
	}

	if qp.Name != "" {
		filter.Name = &qp.Name
	}

	if qp.Price != "" {
		price, err := strconv.Atoi(qp.Price)
		if err != nil {
			return productsvc.Filter{}, validate.NewFieldsError("price", err)
		}
		filter.Price = &price
	}

	if qp.StartCreatedAt != "" {
		t, err := time.Parse(time.RFC3339, qp.StartCreatedAt)
		if err != nil {
			return productsvc.Filter{}, validate.NewFieldsError("start_created_at", err)
		}
		filter.StartCreatedAt = &t
	}

	if qp.EndCreatedAt != "" {
		t, err := time.Parse(time.RFC3339, qp.EndCreatedAt)
		if err != nil {
			return productsvc.Filter{}, validate.NewFieldsError("end_created_at", err)
		}
		filter.EndCreatedAt = &t
	}

	return filter, nil
}
