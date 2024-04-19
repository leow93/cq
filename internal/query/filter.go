package query

import (
	"errors"
	"github.com/leow93/cq/internal/csv"
	"strings"
)

type Operator int

const (
	Eq Operator = iota
	Gt
	Gte
	Lt
	Lte
	Neq
)

var Operators = []string{"!=", ">=", "<=", "=", ">", "<"}

func NewOperator(s string) (Operator, error) {
	switch s {
	case "=":
		return Eq, nil
	case ">":
		return Gt, nil
	case "<":
		return Lt, nil
	case ">=":
		return Gte, nil
	case "<=":
		return Lte, nil
	case "!=":
		return Neq, nil

	default:
		return Eq, errors.New("Unknown operator")
	}
}

func (op Operator) String() string {
	switch op {
	case Eq:
		return "="
	case Gt:
		return ">"
	case Gte:
		return ">="
	case Lt:
		return "<"
	case Lte:
		return "<="
	case Neq:
		return "!="
	default:
		return ""
	}
}

type Filter struct {
	Column   string
	Value    string
	Operator Operator
}

func findOperator(s string) (Operator, error) {
	for _, op := range Operators {
		if strings.Contains(s, op) {
			operator, err := NewOperator(op)
			if err != nil {
				return Eq, err
			}
			return operator, nil
		}

	}
	return Eq, errors.New("Unknown operator")
}

func buildFilter(x string) (*Filter, error) {
	operator, err := findOperator(x)
	if err != nil {
		return nil, err
	}
	parts := strings.Split(x, operator.String())
	if len(parts) != 2 {
		return nil, errors.New("Filter in wrong format")
	}
	column, value := parts[0], parts[1]

	return &Filter{
		Column:   column,
		Value:    value,
		Operator: operator,
	}, nil
}

/*
ParseFilters
Accepts filters of the form "name=bob,age>39"
*/
func ParseFilters(filters string) ([]Filter, error) {
	if filters == "" {
		return nil, nil
	}

	var result []Filter
	filterStrings := strings.Split(filters, ",")
	for _, x := range filterStrings {
		filter, err := buildFilter(x)
		if err == nil {
			result = append(result, *filter)
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
	case Gte:
		return value >= filter.Value
	case Lte:
		return value <= filter.Value
	case Neq:
		return value != filter.Value
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
