package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/caiojorge/goexpert-desafio1/internal/infra"
)

func QuoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		log.Println("get quote")
		q := NewAwesomeApi()
		data, err := q.GetQuoteHandler(w, r)
		if err != nil {
			log.Println("failure - processing the request: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonQuote, err := json.Marshal(data.GetBid())
		if err != nil {
			log.Println("failure - marshaling json: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// poderia usar um service aqui e um repositorio, mas como é um projeto pequeno, não achei necessário
		infra.NewSqlLiteDb().InsertQuote("Dólar", data.GetBid())
		infra.NewFile().Save(data.GetBid())

		w.Write(jsonQuote)

	case http.MethodPost:
		fmt.Fprintf(w, "Handled POST request\n")
	case http.MethodPut:
		fmt.Fprintf(w, "Handled PUT request\n")
	case http.MethodDelete:
		fmt.Fprintf(w, "Handled DELETE request\n")
	default:
		log.Println("failure - unsupported http method")
		http.Error(w, "Unsupported HTTP method", http.StatusMethodNotAllowed)
	}

}
