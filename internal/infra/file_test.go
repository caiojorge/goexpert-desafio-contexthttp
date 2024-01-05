package infra

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

// TestSave tests the Save function
// TestSave tests the Save function
func TestSave(t *testing.T) {
	testValue := 5.25
	Save(testValue) // Assuming Save is part of the same package

	// Open the file to read
	file, err := os.Open("cotacao.txt")
	if err != nil {
		t.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	// Read the last line from the file
	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	// Check if the last line is as expected
	expected := "Dólar: 5.25"
	if !strings.Contains(lastLine, expected) {
		t.Errorf("Expected last line to contain %q, got %q", expected, lastLine)
	}
}
