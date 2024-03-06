package main

import (
	"cq/csv"
	"cq/input"
	"cq/output"
	"fmt"
	"log"
)

func main() {
	data := input.ReadInput()
	err, table := csv.Parser(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output.ToCsv(table))
}
