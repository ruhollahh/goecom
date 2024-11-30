package producthandler

import (
	"github.com/ruhollahh/go-ecom/internal/entity/sort"
	productsvc "github.com/ruhollahh/go-ecom/internal/service/product"
	"github.com/ruhollahh/go-ecom/pkg/validate"
)

var orderByFields = map[string]string{
	"id":    productsvc.OrderByID,
	"name":  productsvc.OrderByName,
	"price": productsvc.OrderByPrice,
}

func parseSort(qp QueryParams) (sort.Sort, error) {
	sortBy, err := sort.Parse(qp.SortField, qp.SortDirection, orderByFields, productsvc.DefaultOrderBy)
	if err != nil {
		return sort.Sort{}, validate.NewFieldsError("sort", err)
	}

	return sortBy, nil
}
