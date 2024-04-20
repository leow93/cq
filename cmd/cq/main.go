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
	sort      query.Sort
}

func chooseFormatter(format *string) output.Formatter {
	switch *format {
	case "json":
		return output.ToJson
	default:
		return output.ToCsv
	}
}

func parseArguments(outputFormat *string, filter *string, sort *string) (*Arguments, error) {
	filters, err := query.ParseFilters(*filter)
	if err != nil {
		return nil, err
	}
	sorter := query.NewSort(*sort)

	return &Arguments{
		formatter: chooseFormatter(outputFormat),
		filter:    filters,
		sort:      sorter,
	}, nil
}

func main() {
	outputFormat := flag.String("output", "csv", "Options are 'csv' (default) and 'json'. e.g. -output=json")
	filter := flag.String("filter", "", "e.g. -filter=age>40,age<=60")
	sort := flag.String("sort", "", "e.g. -sort=age")
	flag.Parse()
	arguments, err := parseArguments(outputFormat, filter, sort)
	if err != nil {
		log.Fatal(err)
	}
	data := input.ReadInput()
	err, table := csv.Parser(data)
	if err != nil {
		log.Fatal(err)
	}
	if arguments.filter != nil {
		table = query.ApplyFilters(arguments.filter, table)
	}
	query.ApplySort(arguments.sort, table)
	formatter := arguments.formatter
	err, output := formatter(table)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(output)
}
