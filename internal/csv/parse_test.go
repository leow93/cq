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
	t.Run("it returns a table with one Column", func(t *testing.T) {
		input := "Name\nbob"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if len(table.Columns) != 1 {
			t.Errorf("Expected 1 Column, got %d", len(table.Columns))
		}
		if len(table.Rows) != 1 {
			t.Errorf("Expected 1 Row, got %d", len(table.Rows))
		}

		expected := map[string]string{"Name": "bob"}
		if !mapEqual(table.Rows[0].Values, expected) {
			t.Errorf("Expected %v, got %v", expected, table.Rows[0].Values)
		}
	})

	t.Run("it returns a multi-Column table", func(t *testing.T) {
		input := "Name,age,email\nbob,30,bob@bob.com\njane,25,jane@jane.com"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		if len(table.Columns) != 3 {
			t.Errorf("Expected 3 Columns, got %d", len(table.Columns))
		}
		if len(table.Rows) != 2 {
			t.Errorf("Expected 2 Rows, got %d", len(table.Rows))
		}
		bob := table.Rows[0].Values
		expectedBob := map[string]string{"Name": "bob", "age": "30", "email": "bob@bob.com"}
		if !mapEqual(bob, expectedBob) {
			t.Errorf("Expected %v, got %v", expectedBob, bob)
		}
		jane := table.Rows[1].Values
		expectedJane := map[string]string{"Name": "jane", "age": "25", "email": "jane@jane.com"}
		if !mapEqual(jane, expectedJane) {
			t.Errorf("Expected %v, got %v", expectedJane, jane)
		}
	})

	t.Run("it handles empty Values", func(t *testing.T) {
		input := "Name,age,email\nbob,30,\njane,25,"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		bob := table.Rows[0].Values
		expectedBob := map[string]string{"Name": "bob", "age": "30", "email": ""}
		if !mapEqual(bob, expectedBob) {
			t.Errorf("Expected %v, got %v", expectedBob, bob)
		}
		jane := table.Rows[1].Values
		expectedJane := map[string]string{"Name": "jane", "age": "25", "email": ""}
		if !mapEqual(jane, expectedJane) {
			t.Errorf("Expected %v, got %v", expectedJane, jane)
		}
	})

	t.Run("it can parse a more complicated table with full sentences in the columns", func(t *testing.T) {
		input := "date,description,severity\n"
		input += "2024-01-01,"
		desc := "\"To be, or not to be, that is the question: Whether 'tis nobler in the mind to suffer The slings and arrows of outrageous fortune, Or to take arms against a sea of troubles And by opposing end them\". - Shakespeare"
		input += desc + ","
		input += "major"
		err, table := Parser(input)
		if err != nil {
			t.Errorf("Expected no error, got %s", err)
		}
		row := table.Rows[0].Values
		expectedRow := map[string]string{"date": "2024-01-01", "description": desc, "severity": "major"}
		if !mapEqual(row, expectedRow) {
			t.Errorf("Expected %v, got %v", expectedRow, row)
		}

	})
}
