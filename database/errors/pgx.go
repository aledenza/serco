package databaseErrors

import (
	"fmt"

	"github.com/jackc/pgx/v5/pgconn"
)

func ParseError(err error) error {
	pgxErr, ok := err.(*pgconn.PgError)
	if !ok {
		return err
	}
	switch pgxErr.Code {
	case "23502":
		return NotNullViolationError{Err: pgxErr}
	case "23505":
		return UniqueViolationError{Err: pgxErr}
	default:
		return err
	}
}

type NotNullViolationError struct{ Err *pgconn.PgError }

func (s NotNullViolationError) Error() string {
	if s.Err == nil {
		return "NotNullViolationError"
	}
	return fmt.Sprintf("NotNullViolationError: %s | %s", s.Err.ColumnName, s.Err.Detail)
}

type UniqueViolationError struct{ Err *pgconn.PgError }

func (s UniqueViolationError) Error() string {
	if s.Err == nil {
		return "UniqueViolationError"
	}
	return fmt.Sprintf(
		"UniqueViolationError: %s | %s",
		s.Err.ConstraintName,
		s.Err.Detail,
	)
}
