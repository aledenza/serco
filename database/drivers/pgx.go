package databaseDrivers

import (
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type pgxDriver struct{}

var Pgx = pgxDriver{}

func (pgxDriver) DriverName() string {
	return "pgx"
}

func (pgxDriver) BuildNamedArgs(args ...map[string]any) any {
	if len(args) == 0 {
		return nil
	}
	return pgx.NamedArgs(args[0])
}
