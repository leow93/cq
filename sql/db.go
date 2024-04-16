package sql

import (
	"cq/csv"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDb() *sql.DB {
	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func CreateTable(db *sql.DB, table csv.Table) {
	statement := "CREATE TABLE " + table.Name + " "
	for i, col := range table.Columns {
		if i == 0 {
			statement += "("
		}
		statement += col.Name + " TEXT"
		if i != len(table.Columns)-1 {
			statement += ", "
		} else {
			statement += ");"
		}
	}
	_, err := db.Exec(statement)
	if err != nil {
		log.Fatal(err)
	}
}
