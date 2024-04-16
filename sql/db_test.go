package sql

import (
	"cq/csv"
	"testing"
)

func TestInitDb(t *testing.T) {
	t.Run("we can initialise a memory database", func(t *testing.T) {
		db := InitDb()
		defer db.Close()

		var number int
		err := db.QueryRow("SELECT 1").Scan(&number)
		if err != nil {
			t.Errorf("Expected no error, got %e", err)
		}
		if number != 1 {
			t.Errorf("Expected 1, got %d", number)
		}
	})
}

func TestCreateTable(t *testing.T) {
	t.Run("it creates a `temp` table", func(t *testing.T) {
		db := InitDb()
		defer db.Close()
		columns := []csv.Column{{Name: "age"}, {Name: "name"}}
		table := csv.Table{
			Columns: columns,
			Name:    "temp",
		}
		CreateTable(db, table)

		var count int
		err := db.QueryRow("SELECT COUNT(*) from temp").Scan(&count)
		if err != nil {
			t.Errorf("Expected no error, got %e", err)
		}
		if count != 0 {
			t.Errorf("Expected 1, got %d", count)
		}
	})
}
