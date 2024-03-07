package query

import (
	"cq/csv"
	"fmt"
	"strings"
)

/**
A query engine for providing SQL-like queries over CSV files.
*/

const (
	SelectAll = "*"
)

type Selection = []string
type WhereClause struct {
	Column string
	Value  string
}

type Query struct {
	Selection    Selection
	WhereClauses []WhereClause
}

const SELECT = "SELECT"
const FROM = "FROM"
const WHERE = "WHERE"
const AND = "AND"

//const OR = "OR"

var keywords []string

func init() {
	keywords = []string{SELECT, FROM, WHERE, AND}
}

func ParseQuery(query string) (error, Query) {
	words := strings.SplitAfter(" ", query)
	fmt.Println(words)
	q := Query{}
	i := 0
	for i < len(words) {
		switch words[i] {
		case SELECT:
			q.Selection = append(q.Selection, words[i+1])
			i += 2
		//case FROM:
		//	q.Table = words[i+1]
		//	i += 2
		case WHERE:
			q.WhereClauses = append(q.WhereClauses, WhereClause{words[i+1], words[i+3]})
			i += 4
		}
		i++
	}

	return nil, q
}

func RunQuery(query string, table csv.Table) (error, csv.Table) {
	return nil, table
}
