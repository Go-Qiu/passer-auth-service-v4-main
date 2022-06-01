package config

import (
	"database/sql"
	"log"
)

// pointer to a database connection.
var db *sql.DB

func init() {
	log.Println("trigger init() in app.go")
}

// Connect will attempt to make make a connection to the database, using the data source name set
// in the .env file.
func Connect() {

	d, err := sql.Open("mysql", "")
	if err != nil {
		panic(err)
	}

	db = d
}

// GetDB will return the active database connection obtain after using Connect().
func GetDB() *sql.DB {
	return db
}
