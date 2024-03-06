package csv

import (
	"testing"
)

func mapEqual(a map[string]string, b map[string]string) bool {
	if len(a) != len(b) {
		return false
	}
	for k, v := range a {
		if b[k] != v {
			return false
		}
	}
	return true
}

func TestParser(t *testing.T) {
	t.Run("it returns a table with one column", func(t *testing.T) {
		input := "name\nbob"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if len(table.columns) != 1 {
			t.Errorf("Expected 1 column, got %d", len(table.columns))
		}
		if len(table.rows) != 1 {
			t.Errorf("Expected 1 row, got %d", len(table.rows))
		}

		expected := map[string]string{"name": "bob"}
		if !mapEqual(table.rows[0].values, expected) {
			t.Errorf("Expected %v, got %v", expected, table.rows[0].values)
		}
	})

	t.Run("it returns a multi-column table", func(t *testing.T) {
		input := "name,age,email\nbob,30,bob@bob.com\njane,25,jane@jane.com"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if len(table.columns) != 3 {
			t.Errorf("Expected 3 columns, got %d", len(table.columns))
		}
		if len(table.rows) != 2 {
			t.Errorf("Expected 2 rows, got %d", len(table.rows))
		}
		bob := table.rows[0].values
		expectedBob := map[string]string{"name": "bob", "age": "30", "email": "bob@bob.com"}
		if !mapEqual(bob, expectedBob) {
			t.Errorf("Expected %v, got %v", expectedBob, bob)
		}
		jane := table.rows[1].values
		expectedJane := map[string]string{"name": "jane", "age": "25", "email": "jane@jane.com"}
		if !mapEqual(jane, expectedJane) {
			t.Errorf("Expected %v, got %v", expectedJane, jane)
		}
	})

	t.Run("it handles empty values", func(t *testing.T) {
		input := "name,age,email\nbob,30,\njane,25,"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		bob := table.rows[0].values
		expectedBob := map[string]string{"name": "bob", "age": "30", "email": ""}
		if !mapEqual(bob, expectedBob) {
			t.Errorf("Expected %v, got %v", expectedBob, bob)
		}
		jane := table.rows[1].values
		expectedJane := map[string]string{"name": "jane", "age": "25", "email": ""}
		if !mapEqual(jane, expectedJane) {
			t.Errorf("Expected %v, got %v", expectedJane, jane)
		}
	})
}
