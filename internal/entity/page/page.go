// Package page provides support for query paging.
package page

import (
	"fmt"
	"strconv"
)

// Page represents the requested page and rows per page.
type Page struct {
	number int
	rows   int
}

// Parse parses the strings and validates the values are in reason.
func Parse(page string, rowsPerPage string) (Page, error) {
	number := 1
	if page != "" {
		var err error
		number, err = strconv.Atoi(page)
		if err != nil {
			return Page{}, fmt.Errorf("page is not a number")
		}
	}

	rows := 10
	if rowsPerPage != "" {
		var err error
		rows, err = strconv.Atoi(rowsPerPage)
		if err != nil {
			return Page{}, fmt.Errorf("rows is not a number")
		}
	}

	if number <= 0 {
		return Page{}, fmt.Errorf("page value too small, must be larger than 0")
	}

	if rows <= 0 {
		return Page{}, fmt.Errorf("rows value too small, must be larger than 0")
	}

	if rows > 100 {
		return Page{}, fmt.Errorf("rows value too large, must be less than 100")
	}

	p := Page{
		number: number,
		rows:   rows,
	}

	return p, nil
}

// String implements the stringer interface.
func (p Page) String() string {
	return fmt.Sprintf("page: %d rows: %d", p.number, p.rows)
}

// Number returns the page number.
func (p Page) Number() int {
	return p.number
}

// RowsPerPage returns the rows per page.
func (p Page) RowsPerPage() int {
	return p.rows
}
