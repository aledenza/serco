package database

import (
	"context"
	"database/sql"
	"errors"

	databaseErrors "github.com/aledenza/serco/database/errors"
	"github.com/aledenza/serco/utils"
	"github.com/jmoiron/sqlx"
)

func (d *Database) FetchVal(
	method string,
) func(ctx context.Context, query string, args ...map[string]any) (any, error) {
	var err error
	metric := databaseMetric(d.name, method)
	return func(ctx context.Context, query string, args ...map[string]any) (any, error) {
		defer func() { metric(err) }()
		var rows *sqlx.Rows
		rows, err = d.conn.NamedQueryContext(ctx, query, utils.GetOptionalValueOrNil(args...))
		if err != nil {
			return nil, databaseErrors.QueryError{Err: err}
		}
		defer rows.Close()
		var result any
		for rows.Next() {
			err = rows.Scan(&result)
			if err != nil {
				return nil, databaseErrors.RowScanError{Err: err}
			}
			break
		}
		return result, err
	}
}

func (d *Database) Execute(method string) func(ctx context.Context, query string, args ...map[string]any) error {
	var err error
	metric := databaseMetric(d.name, method)
	return func(ctx context.Context, query string, args ...map[string]any) error {
		defer func() { metric(err) }()
		_, err = d.conn.NamedExecContext(ctx, query, utils.GetOptionalValueOrNil(args...))
		return err
	}
}

func (d *Database) ExecuteMany(method string) func(ctx context.Context, query string, args []map[string]any) error {
	var err error
	metric := databaseMetric(d.name, method)
	return func(ctx context.Context, query string, args []map[string]any) error {
		defer func() { metric(err) }()
		_, err = d.conn.NamedExecContext(ctx, query, args)
		return err
	}
}

func (d *Database) FetchOne(
	method string,
) func(ctx context.Context, query string, args map[string]any) (map[string]any, error) {
	var err error
	metric := databaseMetric(d.name, method)
	return func(ctx context.Context, query string, args map[string]any) (map[string]any, error) {
		defer func() { metric(err) }()
		var rows *sqlx.Rows
		rows, err = d.conn.NamedQueryContext(ctx, query, args)
		if err != nil {
			return nil, databaseErrors.QueryError{Err: err}
		}
		defer rows.Close()
		resultMap := map[string]any{}
		for rows.Next() {
			err = rows.MapScan(resultMap)
			if err != nil {
				return nil, databaseErrors.RowScanError{Err: err}
			}
			break
		}
		return resultMap, err
	}
}

func (d *Database) FetchAll(
	method string,
) func(ctx context.Context, query string, args map[string]any) ([]map[string]any, error) {
	var err error
	metric := databaseMetric(d.name, method)
	return func(ctx context.Context, query string, args map[string]any) ([]map[string]any, error) {
		defer func() { metric(err) }()
		var rows *sqlx.Rows
		rows, err = d.conn.NamedQueryContext(ctx, query, args)
		if err != nil {
			return nil, databaseErrors.QueryError{Err: err}
		}
		defer rows.Close()
		var result []map[string]any
		for rows.Next() {
			var anyRows []any
			anyRows, err = rows.SliceScan()
			result = make([]map[string]any, 0, len(anyRows))
			for _, row := range anyRows {
				trow, ok := row.(map[string]any)
				if !ok {
					return nil, databaseErrors.RowScanError{}
				}
				result = append(result, trow)
			}
			if err != nil {
				return nil, databaseErrors.RowScanError{Err: err}
			}
		}
		return result, err
	}
}

func (d *Database) Transaction(
	method string,
) func(ctx context.Context, transaction func() error) error {
	var err error
	metric := databaseMetric(d.name, method)
	return func(ctx context.Context, transaction func() error) error {
		defer func() { metric(err) }()
		var tx *sql.Tx
		tx, err = d.conn.BeginTx(ctx, nil)
		if err != nil {
			goto rollback
		}
		err = transaction()
		if err != nil {
			goto rollback
		}
		goto commit

	commit:
		{
			err = tx.Commit()
			return err
		}
	rollback:
		{
			err = errors.Join(err, tx.Rollback())
			return err
		}

	}
}
