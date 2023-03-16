package main

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func connect() (*sql.DB, error) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=ohagi dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
