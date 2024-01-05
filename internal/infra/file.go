package infra

import (
	"fmt"
	"log"
	"os"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (f *File) Save(dollarValue float64) {
	file, err := os.Create("cotacao.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "Dólar: %.2f\n", dollarValue)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Data appended to file successfully")
}
