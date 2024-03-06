package main

import (
	"cq/csv"
	"cq/input"
	"cq/output"
	"fmt"
	"log"
)

func chooseFormatter() output.Formatter {
	return output.ToJson
}

func main() {
	data := input.ReadInput()
	err, table := csv.Parser(data)
	if err != nil {
		log.Fatal(err)
	}
	formatter := chooseFormatter()
	err, output := formatter(table)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output)
}
