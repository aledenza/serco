package database

import (
	"context"
	"time"

	databaseDrivers "github.com/aledenza/serco/database/drivers"
	databaseErrors "github.com/aledenza/serco/database/errors"
	"github.com/aledenza/serco/utils"
	"github.com/creasty/defaults"
	"github.com/jmoiron/sqlx"
)

type Connection struct {
	*sqlx.DB
	driver databaseDrivers.Driver
}

func NewConnection(url string, driver databaseDrivers.Driver, options ...DatabaseOptions) (*Connection, error) {
	db, err := sqlx.Open(string(driver.DriverName()), url)
	if err != nil {
		return nil, databaseErrors.CreateConnectionError{Err: err}
	}
	opts := utils.GetOptionalValueOrNil(options...)
	if opts != nil {
		defaults.Set(opts)
		db.SetConnMaxIdleTime(opts.MaxIdleTime)
		db.SetConnMaxLifetime(opts.MaxLifeTime)
		db.SetMaxIdleConns(opts.MaxIdleConns)
		db.SetMaxOpenConns(opts.MaxOpenConns)
	}
	conn := Connection{DB: db, driver: driver}
	var connErr error
	{
		for range 3 {
			ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
			defer cancel()
			if connErr = conn.PingContext(ctx); connErr != nil {
				time.Sleep(1 * time.Second)
				continue
			}
			return &conn, nil
		}
	}
	return nil, connErr
}

func (c *Connection) Shutdown() {
	if c == nil {
		return
	}
	c.Close()
}
