package query

import (
	"errors"
	"github.com/leow93/cq/internal/csv"
	"strings"
)

type Operator int

const (
	Eq Operator = iota + 1
	Gt
	Gte
	Lt
)

var Operators = []string{"=", ">", "<", ">="}

//
//func (s Operator) String() string {
//	if s < Eq || s > Lt {
//		return fmt.Sprintf("Operator(%d)", int(s))
//	}
//	return Operators[s-1]
//}

func parseOperator(s string) (Operator, error) {
	switch s {
	case "=":
		return Eq, nil
	case ">":
		return Gt, nil
	case "<":
		return Lt, nil
	case ">=":
		return Gte, nil
	default:
		return Eq, errors.New("unknown operator")
	}
}

type Filter struct {
	Column   string
	Value    string
	Operator Operator
}

func tryToken(s string, idx int) (string, string, Operator, error) {
	if idx > len(Operators)-1 {
		return "", "", Eq, errors.New("Unknown operator")
	}

	operator := Operators[idx]
	xs := strings.Split(s, operator)
	switch len(xs) {
	case 2:
		op, err := parseOperator(operator)
		if err != nil {
			return tryToken(s, idx+1)
		}
		return xs[0], xs[1], op, nil
	default:
		return tryToken(s, idx+1)
	}
}

/*
Accepts filters of the form "name=bob,age>39"
*/
func buildFilter(x string) (Filter, error) {
	column, value, op, err := tryToken(x, 0)
	if err != nil {
		return Filter{}, errors.New("Filter in wrong format")
	}
	return Filter{
		Column:   column,
		Value:    value,
		Operator: op,
	}, nil
}

func ParseFilters(filters string) ([]Filter, error) {
	var result []Filter
	filterStrings := strings.Split(filters, ",")
	for _, x := range filterStrings {
		filter, err := buildFilter(x)
		if err == nil {
			result = append(result, filter)
		} else {
			return nil, err
		}
	}
	return result, nil
}

func applyFilter(filter Filter, row csv.Row) bool {
	value, ok := row.Values[filter.Column]
	if !ok {
		return false
	}

	switch filter.Operator {
	case Eq:
		return value == filter.Value
	case Gt:
		return value > filter.Value
	case Lt:
		return value < filter.Value
	default:
		return true
	}
}

func ApplyFilters(filters []Filter, table csv.Table) csv.Table {
	var rows []csv.Row

	for _, row := range table.Rows {
		for _, filter := range filters {
			if applyFilter(filter, row) {
				rows = append(rows, row)
			}
		}
	}
	return csv.NewTable(table.Columns, rows)
}
