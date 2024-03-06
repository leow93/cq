package csv

import (
	"errors"
	"strings"
)

type column struct {
	name string
}
type row struct {
	values map[string]string
}
type Table struct {
	columns []column
	rows    []row
}

func parseColumns(header string) []column {
	columns := strings.Split(header, ",")
	var result []column
	for _, name := range columns {
		result = append(result, column{name: name})
	}
	return result
}

func parseRows(columns []column, rows []string) []row {
	var result []row
	for _, r := range rows {
		rawValues := strings.Split(r, ",")
		values := make(map[string]string, len(columns))
		for i, v := range rawValues {
			columnName := columns[i].name
			values[columnName] = v
		}
		result = append(result, row{values: values})
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
		columns: columns,
		rows:    rows,
	}
}
