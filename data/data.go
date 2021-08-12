package data

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")

	if err != nil{
		return err
	}

	return db.Ping()
}

func CreateStudybuddyDatabase(){
	createNoteDatabaseQuery := `CREATE TABLE IF NOT EXISTS studybuddy(
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	statement, err := db.Prepare(createNoteDatabaseQuery)

	if err != nil{
		log.Fatal(err)
	}

	statement.Exec()

	fmt.Println("studybuddy table created successfully")
}