package output

import (
	"cq/csv"
	"testing"
)

func TestToCsv(t *testing.T) {
	t.Run("it returns a csv string", func(t *testing.T) {
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
		expected := "name,age\nbob,30\njane,25\n"
		_, actual := ToCsv(table)
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})

	t.Run("it handles empty values", func(t *testing.T) {
		table := csv.Table{
			Columns: []csv.Column{
				{Name: "name"},
				{Name: "age"},
			},
			Rows: []csv.Row{
				{Values: map[string]string{"name": "bob", "age": ""}},
				{Values: map[string]string{"name": "jane", "age": "25"}},
			},
		}
		expected := "name,age\nbob,\njane,25\n"
		_, actual := ToCsv(table)
		if actual != expected {
			t.Errorf("Expected %s, got %s", expected, actual)
		}
	})
}
