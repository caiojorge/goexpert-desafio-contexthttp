package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

func Connect() (*sql.DB, error) {
	_db, err := sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("connection ok")
	return _db, nil
}

func Migration(db *sql.DB) {

	sqlStmt := `
    CREATE TABLE example (
        id INTEGER NOT NULL PRIMARY KEY,
        name TEXT
    );`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	log.Println("migration ok")
}

func InsertQuote(db *sql.DB) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	_, err := db.ExecContext(ctx, "INSERT INTO example (id, name) VALUES (?, ?)", 1, "Sample Data")
	if err != nil {
		log.Println(err)
	}
	log.Println("Processing ok")
}

func SelectQuote(db *sql.DB) {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	rows, err := db.QueryContext(ctx, "SELECT id, name FROM example")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
