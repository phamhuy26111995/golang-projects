package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB()  {
	var err error
    DB, err = sql.Open("sqlite3", "api.db")
	if err != nil {
		panic(err)
	}
	
	DB.SetMaxOpenConns(10)
	DB.SetConnMaxIdleTime(5)

	createTables()	
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL,
			password TEXT NOT NULL
	)`


	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic(err)
	}

	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			date_time DATETIME NOT NULL,
			user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
		);`


	_, err = DB.Exec(createEventsTable)

	if(err != nil) {
		panic(err)
	}
}