package output

import "github.com/leow93/cq/internal/csv"

func header(columns []csv.Column) string {
	var result string
	for _, c := range columns {
		result += c.Name + ","
	}
	return result[:len(result)-1] + "\n"
}

func ToCsv(table csv.Table) (error, string) {
	result := header(table.Columns)
	for _, r := range table.Rows {
		for _, c := range table.Columns {
			result += r.Values[c.Name] + ","
		}
		result = result[:len(result)-1] + "\n"
	}
	return nil, result
}
