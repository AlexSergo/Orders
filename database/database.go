package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var database *sql.DB

func Connect() {
	connStr := "user=postgres password=root dbname=first sslmode=disable"
	var err error
	database, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = database.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func Get() *sql.DB {
	return database
}
