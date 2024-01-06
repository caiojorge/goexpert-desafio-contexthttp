package main

import (
	"net/http"

	"github.com/caiojorge/goexpert-desafio1/internal/handler"
)

func main() {
	http.HandleFunc("/quote", handler.QuoteHandler)
	http.HandleFunc("/cotacao", handler.QuoteHandler)
	http.ListenAndServe(":8080", nil)
}
