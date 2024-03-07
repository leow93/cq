package query

import "testing"

func TestParseQuery(t *testing.T) {
	t.Run("simple select all query", func(t *testing.T) {
		_, query := ParseQuery("SELECT * FROM table")
		if len(query.Selection) != 1 {
			t.Errorf("Expected one selection, got %d", len(query.Selection))
		}
		if query.Selection[0] != SelectAll {
			t.Errorf("Expected %s, got %s", SelectAll, query.Selection[0])
		}
	})
}
