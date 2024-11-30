package productdb

import (
	"fmt"
	"github.com/ruhollahh/go-ecom/internal/entity/sort"
	productsvc "github.com/ruhollahh/go-ecom/internal/service/product"
)

var orderByFields = map[string]string{
	productsvc.OrderByID:    "id",
	productsvc.OrderByName:  "name",
	productsvc.OrderByPrice: "price",
}

func orderByClause(sortBy sort.Sort) (string, error) {
	by, exists := orderByFields[sortBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", sortBy.Field)
	}

	return " ORDER BY " + by + " " + sortBy.Direction, nil
}
