package handler

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/caiojorge/goexpert-desafio1/internal/domain"
)

type AwesomeApi struct {
	Quote domain.Quote
}

func NewAwesomeApi() *AwesomeApi {
	return &AwesomeApi{}
}

func (q *AwesomeApi) GetQuoteHandler(w http.ResponseWriter, r *http.Request) (*domain.Quote, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()

	log.Println("Request iniciada")
	defer log.Println("Request finalizada")

	data, err := getQuote(ctx, "https://economia.awesomeapi.com.br/json/last/USD-BRL")
	if err != nil {
		log.Println("failure - internal error: " + err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return data, nil
}

func getQuote(ctx context.Context, url string) (*domain.Quote, error) {
	var quote *domain.Quote
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		log.Println("failure - creating context: " + err.Error())
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("failure - requesting the url: " + err.Error())
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("failure - reading the body: " + err.Error())
		return nil, err
	}

	err = json.Unmarshal(body, &quote)
	if err != nil {
		log.Println("failure - unmarshal json: " + err.Error())
		return nil, err
	}
	return quote, nil
}
