package productdb

import (
	"bytes"
	"context"
	"fmt"
	"github.com/ruhollahh/go-ecom/internal/clients/dbpostgre"
	"github.com/ruhollahh/go-ecom/internal/entity/page"
	"github.com/ruhollahh/go-ecom/internal/entity/product"
	"github.com/ruhollahh/go-ecom/internal/entity/sort"
	productsvc "github.com/ruhollahh/go-ecom/internal/service/product"
)

func (s *Store) Query(ctx context.Context, filter productsvc.Filter, sortBy sort.Sort, page page.Page) ([]product.Product, error) {
	data := map[string]any{
		"offset":        (page.Number() - 1) * page.RowsPerPage(),
		"rows_per_page": page.RowsPerPage(),
	}

	const q = `
    SELECT
	    id, name, description, image_id, price, created_at, updated_at
	FROM
	  	products`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	order, err := orderByClause(sortBy)
	if err != nil {
		return nil, err
	}

	buf.WriteString(order)
	buf.WriteString(" OFFSET :offset ROWS FETCH NEXT :rows_per_page ROWS ONLY")

	var dbPrds []dbProduct
	if err := dbpostgre.NamedQuerySlice(ctx, s.log, s.db, buf.String(), data, &dbPrds); err != nil {
		return nil, fmt.Errorf("namedqueryslice: %w", err)
	}

	prds := toProducts(dbPrds)

	return prds, nil
}

// Count returns the total number of homes in the DB.
func (s *Store) Count(ctx context.Context, filter productsvc.Filter) (int, error) {
	data := map[string]any{}

	const q = `
    SELECT
        count(1)
    FROM
        products`

	buf := bytes.NewBufferString(q)
	s.applyFilter(filter, data, buf)

	var count struct {
		Count int `db:"count"`
	}
	if err := dbpostgre.NamedQueryStruct(ctx, s.log, s.db, buf.String(), data, &count); err != nil {
		return 0, fmt.Errorf("db: %w", err)
	}

	return count.Count, nil
}
