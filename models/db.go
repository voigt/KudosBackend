package models

import (
	"database/sql"
	"log"
	"os"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	// Ping verifies a connection to the database is still alive,
	// establishing a connection if necessary.
	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}

func ResetDB() {
	// Remove old DB
	os.Remove("./kudos.db")
	PopulateDB()
}

func PopulateDB() {

	// creates table
	sqlStmt := `
	create table kudos (id integer not null primary key autoincrement, name text);
	`

	_, err := db.Exec(sqlStmt)
	checkErr(err)
}
