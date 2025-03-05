package databaseErrors

import "fmt"

type NoConnectionError struct{}

func (NoConnectionError) Error() string {
	return "no connection error"
}

type PingError struct{ Err error }

func (PingError) Error() string {
	return "database ping error"
}

type CreateConnectionError struct{ Err error }

func (e CreateConnectionError) Error() string {
	return fmt.Sprintf("create connection error: %s", e.Err)
}

type QueryError struct{ Err error }

func (e QueryError) Error() string {
	return fmt.Sprintf("query error: %s", e.Err)
}

type RowScanError struct{ Err error }

func (e RowScanError) Error() string {
	return fmt.Sprintf("row scan error: %s", e.Err)
}
