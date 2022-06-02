package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func OpenDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// ok. connections pool instantiated.
	// check if the network connection to the database is ok.
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// ok.
	return db, nil
}
