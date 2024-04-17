package output

import "github.com/leow93/cq/internal/csv"

type Formatter = func(table csv.Table) (error, string)
