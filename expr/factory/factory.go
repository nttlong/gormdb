package factory

import (
	"github.com/nttlong/gormdb/expr"

	"github.com/nttlong/gormdb/expr/exprpostgres"
)

func NewExpr(driver string) expr.IExpr {
	switch driver {
	case "postgres":
		return exprpostgres.New()
	default:
		panic("Unsupported driver: " + driver)
	}
}
