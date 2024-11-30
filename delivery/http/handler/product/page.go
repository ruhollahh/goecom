package producthandler

import (
	"github.com/ruhollahh/go-ecom/internal/entity/page"
	"github.com/ruhollahh/go-ecom/pkg/validate"
)

func parsePage(qp QueryParams) (page.Page, error) {
	pagination, err := page.Parse(qp.Page, qp.Rows)
	if err != nil {
		return page.Page{}, validate.NewFieldsError("page", err)
	}

	return pagination, nil
}
