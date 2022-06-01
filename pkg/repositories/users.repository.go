package repositories

import (
	"database/sql"
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// New returns a new database connection to a MySQL database, uisng the Data Source Name (DSN) passed in.
func New(dsn string) (*sql.DB, error) {

	// instantiate a (go managed) database connections pool.
	// a successful instantiation DOES NOT mean a successful connection to the database (via the network).
	d, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, errors.New("fail to instantiate a database connections pool")
	}

	// connections pool struct is ready.
	// setup the connection pool attributes

	// max. number of open connections
	d.SetMaxOpenConns(10)

	// max. number of idle connections
	d.SetMaxIdleConns(3)

	// max. life time of each connection (since it was opened)
	d.SetConnMaxLifetime(time.Hour)

	// test the connection to the database.
	err = d.Ping()
	if err != nil {
		return nil, errors.New("fail to connect to the database")
	}

	// ok.
	return d, nil
}
