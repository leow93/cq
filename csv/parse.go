package csv

import (
	"errors"
	"strings"
)

type Column struct {
	Name string
}
type Row struct {
	Values map[string]string
}
type Table struct {
	Columns []Column
	Rows    []Row
}

func parseColumns(header string) []Column {
	columns := strings.Split(header, ",")
	var result []Column
	for _, name := range columns {
		result = append(result, Column{Name: name})
	}
	return result
}

func parseRows(columns []Column, rows []string) []Row {
	var result []Row
	for _, r := range rows {
		rawValues := strings.Split(r, ",")
		values := make(map[string]string, len(columns))
		for i, v := range rawValues {
			columnName := columns[i].Name
			values[columnName] = v
		}
		result = append(result, Row{Values: values})
	}
	return result
}

func Parser(input string) (error, Table) {
	lines := strings.Split(input, "\n")

	if len(lines) == 0 {
		return errors.New("No input"), Table{}
	}

	header := lines[0]
	columns := parseColumns(header)
	rows := parseRows(columns, lines[1:])

	return nil, Table{
		Columns: columns,
		Rows:    rows,
	}
}
