package infra

import (
	"testing"
)

func TestConnect(t *testing.T) {
	s := NewSqlLiteDb()
	db, err := s.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %s", err)
	}
	if db == nil {
		t.Fatal("Received nil database connection")
	}
	_ = db.Close()
}

func TestMigration(t *testing.T) {
	s := NewSqlLiteDb()
	db, _ := s.Connect()
	defer db.Close()

	s.Migration()
}

func TestInsertQuote(t *testing.T) {
	s := NewSqlLiteDb()
	db, _ := s.Connect()
	defer db.Close()

	s.Migration()

	err := s.InsertQuote("Dólar", 5.25)
	if err != nil {
		t.Fatalf("Failed to insert quote: %s", err)
	}

	err = s.InsertQuote("Dólar", 5.26)
	if err != nil {
		t.Fatalf("Failed to insert quote: %s", err)
	}

	s.InsertQuote("Dólar", 5.27)
	if err != nil {
		t.Fatalf("Failed to insert quote: %s", err)
	}

	s.SelectQuote()
}
