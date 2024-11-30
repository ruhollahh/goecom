package productsvc

import (
	"context"
	"fmt"
	"github.com/ruhollahh/go-ecom/internal/entity/page"
	"github.com/ruhollahh/go-ecom/internal/entity/product"
	"github.com/ruhollahh/go-ecom/internal/entity/sort"
)

func (s *Service) Query(ctx context.Context, filter Filter, orderBy sort.Sort, page page.Page) ([]product.Product, error) {
	prds, err := s.storer.Query(ctx, filter, orderBy, page)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}

	return prds, nil
}

func (s *Service) Count(ctx context.Context, filter Filter) (int, error) {
	c, err := s.storer.Count(ctx, filter)
	if err != nil {
		return 0, fmt.Errorf("count: %w", err)
	}

	return c, err
}
