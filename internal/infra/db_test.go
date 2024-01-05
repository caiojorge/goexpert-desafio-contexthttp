package infra

import (
	"testing"
)

func TestConnect(t *testing.T) {
	db, err := Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	if db == nil {
		t.Fatal("Received nil database connection")
	}
	_ = db.Close()
}

func TestMigration(t *testing.T) {
	db, _ := Connect()
	defer db.Close()

	Migration(db)
}

func TestInsertQuote(t *testing.T) {
	db, _ := Connect()
	defer db.Close()

	Migration(db)
	InsertQuote(db)
	SelectQuote(db)
}
