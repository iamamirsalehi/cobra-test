package data

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
)

var db *sql.DB

func OpenDatabase() error {
	var err error

	db, err = sql.Open("sqlite3", "./sqlite-database.db")

	if err != nil {
		return err
	}

	return db.Ping()
}

func CreateStudybuddyDatabase() {
	createNoteDatabaseQuery := `CREATE TABLE IF NOT EXISTS studybuddy(
		"id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		"word" TEXT,
		"definition" TEXT,
		"category" TEXT
	);`

	statement, err := db.Prepare(createNoteDatabaseQuery)

	if err != nil {
		log.Fatal(err)
	}

	statement.Exec()

	fmt.Println("studybuddy table created successfully")
}

func InsertNote(word string, definition string, category string) {
	insertQuery := `INSERT INTO studybuddy (word, definition, category) 
					VALUES (?, ?, ?)`

	statement, err := db.Prepare(insertQuery)

	if err != nil {
		log.Fatal(err)
	}

	_, err = statement.Exec(word, definition, category)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("New note successfully created")
}

func DisplayAllNotes() {
	selectQuery := `SELECT * FROM studybuddy`

	rows, err := db.Query(selectQuery)

	if err != nil {
		fmt.Printf("Diplaying all notes failed: %s", err)
		os.Exit(1)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var word string
		var definition string
		var category string

		rows.Scan(&id, &word, &definition, &category)

		log.Println("[", id, "]", "word: ", word, " - definition: ", definition, " - category: ", category)
	}
}
