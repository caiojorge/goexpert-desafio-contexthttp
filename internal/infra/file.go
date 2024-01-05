package infra

import (
	"fmt"
	"log"
	"os"
)

func Save(dollarValue float64) {
	// Replace this with the actual value
	//dollarValue := 5.25 // Example value

	// Open the file in append mode, create it if it does not exist
	//file, err := os.OpenFile("cotacao.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Append formatted content to the file
	_, err = fmt.Fprintf(file, "Dólar: %.2f\n", dollarValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data appended to file successfully")
}
