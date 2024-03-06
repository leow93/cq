package output

import (
	"cq/csv"
	"testing"
)

func TestToJson(t *testing.T) {
	t.Run("it returns a json string", func(t *testing.T) {
		table := csv.Table{
			Columns: []csv.Column{
				{Name: "name"},
				{Name: "age"},
			},
			Rows: []csv.Row{
				{Values: map[string]string{"name": "bob", "age": "30"}},
				{Values: map[string]string{"name": "jane", "age": "25"}},
			},
		}
		expected := `[{"age":"30","name":"bob"},{"age":"25","name":"jane"}]`
		err, actual := ToJson(table)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
}
