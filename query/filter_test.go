package query

import "testing"

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
		result := ParseFilters(input)
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
		result := ParseFilters(input)
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
}
