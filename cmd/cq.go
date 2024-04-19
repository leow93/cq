package main

import (
	"flag"
	"fmt"
	"github.com/leow93/cq/internal/csv"
	"github.com/leow93/cq/internal/input"
	"github.com/leow93/cq/internal/output"
	"github.com/leow93/cq/internal/query"
	"log"
)

type Arguments struct {
	formatter output.Formatter
	filter    []query.Filter
}

func chooseFormatter(format *string) output.Formatter {
	switch *format {
	case "json":
		return output.ToJson
	default:
		return output.ToCsv
	}
}

func parseArguments(outputFormat *string, filter *string) (*Arguments, error) {
	filters, err := query.ParseFilters(*filter)
	if err != nil {
		return nil, err
	}

	return &Arguments{
		formatter: chooseFormatter(outputFormat),
		filter:    filters,
	}, nil
}

func main() {
	outputFormat := flag.String("output", "csv", "Options are 'csv' (default) and 'json'. e.g. -output=json")
	filter := flag.String("filter", "", "e.g. -filter=column1=value1,column2=value2")
	flag.Parse()
	arguments, err := parseArguments(outputFormat, filter)
	if err != nil {
		log.Fatal(err)
	}
	data := input.ReadInput()
	err, table := csv.Parser(data)
	if err != nil {
		log.Fatal(err)
	}
	table = query.ApplyFilters(arguments.filter, table)
	formatter := arguments.formatter
	err, output := formatter(table)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output)
}
