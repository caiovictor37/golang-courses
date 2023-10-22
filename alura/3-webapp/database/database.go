package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectIntoDatabase() *sql.DB {
	// Basic implementation, in PRD password should use environment variables
	connection := "user=alura_store dbname=alura_store password=ExamplePsw321. host=localhost sslmode=disable"
	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}
