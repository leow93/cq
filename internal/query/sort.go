package query

import (
	"github.com/leow93/cq/internal/csv"
	"sort"
)

type Sort struct {
	Column string
}

func NewSort(column string) Sort {
	return Sort{Column: column}
}

// todo: parse sort

func ApplySort(sorter Sort, table csv.Table) {
	sort.SliceStable(table.Rows, func(i, j int) bool {
		a, ok := table.Rows[i].Values[sorter.Column]
		if !ok {
			return false
		}
		b, ok := table.Rows[j].Values[sorter.Column]
		if !ok {
			return true
		}

		return a < b
	})
}
