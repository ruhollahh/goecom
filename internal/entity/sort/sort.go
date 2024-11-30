// Package order provides support for describing the ordering of data.
package sort

import (
	"fmt"
)

const (
	ASC  = "asc"
	DESC = "desc"
)

var directions = map[string]string{
	ASC:  "asc",
	DESC: "desc",
}

type Sort struct {
	Field     string
	Direction string
}

func New(field string, direction string) Sort {
	if _, exists := directions[direction]; !exists {
		return Sort{
			Field:     field,
			Direction: ASC,
		}
	}

	return Sort{
		Field:     field,
		Direction: direction,
	}
}

func Parse(field string, direction string, fieldMappings map[string]string, defaultOrder Sort) (Sort, error) {
	fieldValue := defaultOrder.Field
	directionValue := defaultOrder.Direction
	var exists bool

	if field != "" {
		fieldValue, exists = fieldMappings[field]
		if !exists {
			return Sort{}, fmt.Errorf("unknown order: %s", field)
		}
	}

	if direction != "" {
		directionValue, exists = directions[direction]
		if !exists {
			return Sort{}, fmt.Errorf("unknown direction: %s", direction)
		}
	}

	return New(fieldValue, directionValue), nil
}
