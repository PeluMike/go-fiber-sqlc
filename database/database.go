package database

import (
	"database/sql"
	"log"

	"github.com/PeluMike/blog/src/sqlc"
)

var Queries *sqlc.Queries

func ConnectDb() {
	connStr := "postgres://PeluMike@localhost:5430/go-test?sslmode=disable"
	dbConn, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error connecting db")
	}

	Queries = sqlc.New(dbConn)
}
