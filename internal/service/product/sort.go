package productsvc

import "github.com/ruhollahh/go-ecom/internal/entity/sort"

var DefaultOrderBy = sort.New(OrderByID, sort.ASC)

const (
	OrderByID    = "id"
	OrderByName  = "name"
	OrderByPrice = "price"
)
