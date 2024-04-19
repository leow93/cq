package csv

import (
	"errors"
	"strings"
)

type Column struct {
	Name string
}

func NewColumn(name string) Column {
	return Column{Name: name}
}

type Row struct {
	Values map[string]string
}

func NewRow(values map[string]string) Row {
	return Row{Values: values}
}

type Table struct {
	Columns []Column
	Rows    []Row
}

func NewTable(cols []Column, rows []Row) Table {
	return Table{Columns: cols, Rows: rows}
}

func parseColumns(header string) []Column {
	columns := strings.Split(header, ",")
	var result []Column
	for _, name := range columns {
		result = append(result, Column{Name: name})
	}
	return result
}

func getRawValues(row string) []string {
	var result []string
	temp := ""
	ignoreCommaSeparation := false

	for _, r := range []rune(row) {
		if r == '"' {
			ignoreCommaSeparation = !ignoreCommaSeparation
		}

		if r == ',' {
			if ignoreCommaSeparation {
				temp += ","
			} else {
				result = append(result, temp)
				temp = ""
			}
		} else {
			temp += string(r)
		}
	}

	result = append(result, temp)

	return result
}

func parseRows(columns []Column, rows []string) []Row {
	var result []Row

	for _, r := range rows {
		rawValues := getRawValues(r)
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
		return errors.New("No input"), NewTable(nil, nil)
	}

	header := lines[0]
	columns := parseColumns(header)
	rows := parseRows(columns, lines[1:])

	return nil, NewTable(columns, rows)
}
