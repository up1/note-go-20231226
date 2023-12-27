package demo

import (
	"database/sql"
)

// Connect to postgres database

func CreateConnection() (*sql.DB, error) {
	// connect to postgres db
	db, err := sql.Open("postgres", "host=localhost port=5432 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	// check db
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
