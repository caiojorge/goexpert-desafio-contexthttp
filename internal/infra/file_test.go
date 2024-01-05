package infra

import (
	"bufio"
	"os"
	"strings"
	"testing"
)

func TestSave(t *testing.T) {
	testValue := 5.25
	f := NewFile()

	f.Save(testValue)

	file, err := os.Open("cotacao.txt")
	if err != nil {
		t.Fatalf("Failed to open file: %s", err)
	}
	defer file.Close()

	var lastLine string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lastLine = scanner.Text()
	}

	expected := "Dólar: 5.25"
	if !strings.Contains(lastLine, expected) {
		t.Errorf("Expected last line to contain %q, got %q", expected, lastLine)
	}
}
