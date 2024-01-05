package infra

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "modernc.org/sqlite"
)

type SqlLiteDb struct {
	db *sql.DB
}

func NewSqlLiteDb() *SqlLiteDb {
	return &SqlLiteDb{}
}

func (s *SqlLiteDb) Connect() (*sql.DB, error) {
	var err error
	s.db, err = sql.Open("sqlite", ":memory:")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	log.Println("connection ok")
	return s.db, nil
}

func (s *SqlLiteDb) Migration() {
	if s.db == nil {
		log.Println("database not connected")
		return
	}

	sqlStmt := `
    CREATE TABLE currency (
		id INTEGER NOT NULL PRIMARY KEY,
		name TEXT,
		currency_value FLOAT
	);
	`
	_, err := s.db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("%q: %s\n", err, sqlStmt)
	}
	log.Println("migration ok")
}

func (s *SqlLiteDb) InsertQuote(name string, currencyValue float64) error {

	if s.db == nil {
		log.Println("database not connected")
		return fmt.Errorf("database not connected")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	sql := "INSERT INTO currency (name, currency_value) VALUES (?, ?)"
	_, err := s.db.ExecContext(ctx, sql, name, currencyValue)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Processing ok")
	return nil
}

func (s *SqlLiteDb) SelectQuote() {
	if s.db == nil {
		log.Println("database not connected")
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()

	rows, err := s.db.QueryContext(ctx, "SELECT id, name, currency_value FROM currency")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var currency_value float64
		err = rows.Scan(&id, &name, &currency_value)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name, currency_value)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

}
