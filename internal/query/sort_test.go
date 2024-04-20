package query

import (
	"github.com/leow93/cq/internal/csv"
	"testing"
)

func TestApplySort(t *testing.T) {
	t.Run("it sorts the table in-place, in ascending order", func(t *testing.T) {
		columns := []csv.Column{
			csv.NewColumn("name"),
			csv.NewColumn("age"),
		}
		rows := []csv.Row{
			csv.NewRow(map[string]string{"name": "bob", "age": "50"}),
			csv.NewRow(map[string]string{"name": "alice", "age": "25"}),
		}

		table := csv.Table{
			columns,
			rows,
		}
		sort := NewSort("age")

		ApplySort(sort, table)

		if table.Rows[0].Values["name"] != "alice" {
			t.Errorf("Expected first row to be alice, got %v", table.Rows[0].Values["name"])
		}
		if table.Rows[1].Values["name"] != "bob" {
			t.Errorf("Expected second row to be bob, got %v", table.Rows[1].Values["name"])
		}
	})

	t.Run("it maintains original order of equivalent items", func(t *testing.T) {
		columns := []csv.Column{
			csv.NewColumn("name"),
			csv.NewColumn("age"),
		}
		rows := []csv.Row{
			csv.NewRow(map[string]string{"name": "bob", "age": "50"}),
			csv.NewRow(map[string]string{"name": "alice", "age": "25"}),
			csv.NewRow(map[string]string{"name": "cameron", "age": "15"}),
			csv.NewRow(map[string]string{"name": "alice", "age": "39"}),
		}

		table := csv.Table{
			columns,
			rows,
		}
		sort := NewSort("name")

		ApplySort(sort, table)
		if table.Rows[0].Values["age"] != "25" {
			t.Errorf("Expected first row to be alice (25), got %v", table.Rows[0].Values["name"])
		}
		if table.Rows[1].Values["age"] != "39" {
			t.Errorf("Expected second row to be alice (39), got %v", table.Rows[1].Values["name"])
		}
		if table.Rows[2].Values["age"] != "50" {
			t.Errorf("Expected third row to be bob, got %v", table.Rows[2].Values["name"])
		}
		if table.Rows[3].Values["age"] != "15" {
			t.Errorf("Expected fourth row to be cameron, got %v", table.Rows[2].Values["name"])
		}
	})

	t.Run("sorting has no effect if the sort by is not a column", func(t *testing.T) {
		columns := []csv.Column{
			csv.NewColumn("name"),
			csv.NewColumn("age"),
		}
		rows := []csv.Row{
			csv.NewRow(map[string]string{"name": "bob", "age": "50"}),
			csv.NewRow(map[string]string{"name": "alice", "age": "25"}),
			csv.NewRow(map[string]string{"name": "cameron", "age": "15"}),
			csv.NewRow(map[string]string{"name": "alice", "age": "39"}),
		}

		table := csv.Table{
			columns,
			rows,
		}
		sort := NewSort("foo")
		ApplySort(sort, table)
		if table.Rows[0].Values["name"] != "bob" {
			t.Errorf("Expected row to be bob, got %v", table.Rows[0].Values["name"])
		}
		if table.Rows[1].Values["name"] != "alice" {
			t.Errorf("Expected row to be alice, got %v", table.Rows[0].Values["name"])
		}
		if table.Rows[2].Values["name"] != "cameron" {
			t.Errorf("Expected row to be cameron, got %v", table.Rows[0].Values["name"])
		}
		if table.Rows[3].Values["name"] != "alice" {
			t.Errorf("Expected row to be alice, got %v", table.Rows[0].Values["name"])
		}

	})
}
