package productsvc

import (
	"context"
	"github.com/ruhollahh/go-ecom/internal/entity/page"
	"github.com/ruhollahh/go-ecom/internal/entity/product"
	"github.com/ruhollahh/go-ecom/internal/entity/sort"
	"github.com/ruhollahh/go-ecom/pkg/logger"
)

type Storer interface {
	Query(ctx context.Context, filter Filter, sortBy sort.Sort, page page.Page) ([]product.Product, error)
	Count(ctx context.Context, filter Filter) (int, error)
}

type Service struct {
	storer Storer
	log    *logger.Logger
}

func New(storer Storer) *Service {
	return &Service{
		storer: storer,
	}
}
