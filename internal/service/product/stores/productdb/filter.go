package productdb

import (
	"bytes"
	productsvc "github.com/ruhollahh/go-ecom/internal/service/product"
	"strings"
)

func (s *Store) applyFilter(filter productsvc.Filter, data map[string]any, buf *bytes.Buffer) {
	var wc []string

	if filter.ID != nil {
		data["id"] = *filter.ID
		wc = append(wc, "id = :id")
	}

	if filter.Name != nil {
		data["name"] = filter.Name
		wc = append(wc, "name = :name")
	}

	if filter.Price != nil {
		data["price"] = filter.Price
		wc = append(wc, "price = :price")
	}

	if filter.StartCreatedAt != nil {
		data["start_created_at"] = filter.StartCreatedAt.UTC()
		wc = append(wc, "created_at >= :start_created_at")
	}

	if filter.EndCreatedAt != nil {
		data["end_created_at"] = filter.EndCreatedAt.UTC()
		wc = append(wc, "created_at <= :end_created_at")
	}

	if len(wc) > 0 {
		buf.WriteString(" WHERE ")
		buf.WriteString(strings.Join(wc, " AND "))
	}
}
