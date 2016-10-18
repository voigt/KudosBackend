package models

import (
	"database/sql"
	"os"
)

var db *sql.DB

func InitDB(dataSourceName string) {
	db, err := sql.Open("sqlite3", dataSourceName)
	checkErr(err)

	if db == nil {
		panic("db nil")
	}
}

func PopulateDB() {
	sql := `
		CREATE TABLE IF NOT EXISTS kudos (
			id integer not null primary key autoincrement,
			url text,
			kudos integer
		);
	`

	_, err := db.Exec(sql)
	checkErr(err)
}

func ResetDB() {
	// Remove old DB
	os.Remove("./kudos.db")
	PopulateDB()
}
