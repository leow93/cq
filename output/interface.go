package output

import "cq/csv"

type Formatter = func(table csv.Table) (error, string)
