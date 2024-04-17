package main

import (
	"flag"
	"fmt"
	"github.com/leow93/cq/internal/csv"
	"github.com/leow93/cq/internal/input"
	"github.com/leow93/cq/internal/output"
	"log"
)

type Arguments struct {
	formatter output.Formatter
}

func chooseFormatter(format *string) output.Formatter {
	switch *format {
	case "json":
		return output.ToJson
	default:
		return output.ToCsv
	}
}

func parseArguments(outputFormat *string) Arguments {
	return Arguments{
		formatter: chooseFormatter(outputFormat),
	}
}

func main() {
	outputFormat := flag.String("output", "csv", "Options are 'csv' (default) and 'json'. e.g. -output=json")
	flag.Parse()

	arguments := parseArguments(outputFormat)
	data := input.ReadInput()
	err, table := csv.Parser(data)
	if err != nil {
		log.Fatal(err)
	}
	formatter := arguments.formatter
	err, output := formatter(table)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output)
}
