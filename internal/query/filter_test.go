package query

import (
	"github.com/leow93/cq/internal/csv"
	"testing"
)

func equalFilter(a Filter, b Filter) bool {
	return a.Value == b.Value && a.Operator == b.Operator && a.Column == b.Column
}

func testEquality(t *testing.T, expected Filter, actual Filter) {
	if !equalFilter(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func TestParseFilters(t *testing.T) {

	t.Run("it parses equality filters", func(t *testing.T) {
		input := "name=bob,age=30"
		result, err := ParseFilters(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("Expected 2 filters, got %d", len(result))
		}
		got := result[0]
		expected := Filter{
			Column:   "name",
			Value:    "bob",
			Operator: Eq,
		}
		testEquality(t, expected, got)
		got = result[1]
		expected = Filter{
			Column:   "age",
			Value:    "30",
			Operator: Eq,
		}
		testEquality(t, expected, got)
	})

	t.Run("it parses gt filters", func(t *testing.T) {
		input := "name=bob,age>30"
		result, err := ParseFilters(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("Expected 2 filters, got %d", len(result))
		}
		got := result[1]
		expected := Filter{
			Column:   "age",
			Value:    "30",
			Operator: Gt,
		}
		testEquality(t, expected, got)
	})

	t.Run("it parses lt filters", func(t *testing.T) {
		input := "name=bob,age<30"
		result, err := ParseFilters(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("Expected 2 filters, got %d", len(result))
		}
		got := result[1]
		expected := Filter{
			Column:   "age",
			Value:    "30",
			Operator: Lt,
		}
		testEquality(t, expected, got)
	})

	t.Run("it parses gte filters", func(t *testing.T) {
		input := "age>=30"
		result, err := ParseFilters(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Expected 1 filter, got %d", len(result))
		}
		got := result[0]
		expected := Filter{
			Column:   "age",
			Value:    "30",
			Operator: Gte,
		}
		testEquality(t, expected, got)
	})

	t.Run("it parses lte filters", func(t *testing.T) {
		input := "age<=30"
		result, err := ParseFilters(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Expected 1 filter, got %d", len(result))
		}
		got := result[0]
		expected := Filter{
			Column:   "age",
			Value:    "30",
			Operator: Lte,
		}
		testEquality(t, expected, got)
	})

	t.Run("it parses neq filters", func(t *testing.T) {
		input := "age!=30"
		result, err := ParseFilters(input)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if len(result) != 1 {
			t.Errorf("Expected 1 filter, got %d", len(result))
		}
		got := result[0]
		expected := Filter{
			Column:   "age",
			Value:    "30",
			Operator: Neq,
		}
		testEquality(t, expected, got)
	})
}

func TestApplyFilters(t *testing.T) {
	columns := []csv.Column{
		csv.NewColumn("name"),
		csv.NewColumn("age"),
	}
	rows := []csv.Row{
		csv.NewRow(map[string]string{"name": "bob", "age": "30"}),
		csv.NewRow(map[string]string{"name": "alice", "age": "40"}),
	}
	table := csv.NewTable(columns, rows)
	tests := []struct {
		name           string
		filter         string
		expectedPerson string
	}{
		{
			name:           "equality filter",
			filter:         "name=alice",
			expectedPerson: "alice",
		},
		{
			name:           "greater than filter",
			filter:         "age>35",
			expectedPerson: "alice",
		},
		{
			name:           "less than filter",
			filter:         "age<35",
			expectedPerson: "bob",
		},
		{
			name:           "greater than or equal filter",
			filter:         "age>=40",
			expectedPerson: "alice",
		},
		{
			name:           "less than or equal filter",
			filter:         "age<=30",
			expectedPerson: "bob",
		},
		{
			name:           "not equal filter",
			filter:         "age!=30",
			expectedPerson: "alice",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filters, err := ParseFilters(tt.filter)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			result := ApplyFilters(filters, table)
			if len(result.Rows) != 1 {
				t.Errorf("Expected 1 row, got %d", len(result.Rows))
			}
			if result.Rows[0].Values["name"] != tt.expectedPerson {
				t.Errorf("Expected '%s', got %s", tt.expectedPerson, result.Rows[0].Values["name"])
			}
		})
	}
}
