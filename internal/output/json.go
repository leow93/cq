package output

import (
	"encoding/json"
	"github.com/leow93/cq/internal/csv"
)

func ToJson(table csv.Table) (error, string) {
	rows := table.Rows

	values := make([]map[string]string, len(rows))
	for i, r := range rows {
		values[i] = r.Values
	}

	jsonContent, err := json.Marshal(values)
	if err != nil {
		return err, ""
	}
	return nil, string(jsonContent)
}
