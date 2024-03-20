package query

import (
	"errors"
	"fmt"
	"strings"
)

type Operator int

const (
	Eq Operator = iota + 1
	Gt
)

var Operators = []string{"=", ">"}

func (s Operator) String() string {
	if s < Eq || s > Gt {
		return fmt.Sprintf("Operator(%d)", int(s))
	}
	return Operators[s-1]
}

func (s Operator) IsValid() bool {
	switch s {
	case Eq:
		return true
	}
	return false
}

func parseOperator(s string) (Operator, error) {
	switch s {
	case "=":
		return Eq, nil
	case ">":
		return Gt, nil
	default:
		return Eq, errors.New("unknown operator")
	}
}

type Filter struct {
	Column   string
	Value    string
	Operator Operator
}

func trimEach(xs []string) []string {
	result := make([]string, len(xs))
	for _, x := range xs {
		trimmed := strings.TrimSpace(x)
		if len(trimmed) > 0 {
			result = append(result, strings.TrimSpace(x))
		}
	}
	return result
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

func ParseFilters(filters string) []Filter {
	var result []Filter
	filterStrings := strings.Split(filters, ",")
	//xs := trimEach(filterStrings)
	for _, x := range filterStrings {
		filter, err := buildFilter(x)
		if err == nil {
			result = append(result, filter)
		}
	}
	return result
}
