package repository

import (
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

func GetDatabaseConnection() *sql.DB{
	connString := "postgres://postgres:qburst@localhost/Student"

	db, err := sql.Open("postgres", connString)
	if err!=nil {
		log.Fatal(err)
	}
	return db

}

