package main

import (
	"cq/csv"
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

func RunQuery(query string, table csv.Table) (error, csv.Table) {
	return nil, table
}
