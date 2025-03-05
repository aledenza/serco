package database

import (
	"database/sql"
)

type Database struct {
	conn *Connection
	name string
}

func NewDatabase(name string, conn *Connection) Database {
	return Database{conn: conn, name: name}
}

func prepareNamedArgs(args ...map[string]any) []any {
	var namedArgs []any
	var arg map[string]any
	if len(args) != 0 {
		arg = args[0]
		namedArgs = make([]any, 0, len(arg))
	}
	for k, v := range arg {
		namedArgs = append(namedArgs, sql.Named(k, v))
	}
	return namedArgs
}
