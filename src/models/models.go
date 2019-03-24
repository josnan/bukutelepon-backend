package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

// Init initialize
func Init(connection string) {
	var err error
	db, err = sql.Open("postgres", connection)
	if err != nil {
		log.Fatal("DB Connection Error: ", err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal("DB Error: ", err)
	}
}
