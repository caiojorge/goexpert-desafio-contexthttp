package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type Quote struct {
	USDBRL struct {
		Code       string `json:"code"`
		Codein     string `json:"codein"`
		Name       string `json:"name"`
		High       string `json:"high"`
		Low        string `json:"low"`
		VarBid     string `json:"varBid"`
		PctChange  string `json:"pctChange"`
		Bid        string `json:"bid"`
		Ask        string `json:"ask"`
		Timestamp  string `json:"timestamp"`
		CreateDate string `json:"create_date"`
	} `json:"USDBRL"`
}

func main() {
	http.HandleFunc("/quote", quoteHandler)
	http.HandleFunc("/cotacao", quoteHandler)
	http.ListenAndServe(":8080", nil)
}

func quoteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		log.Println("get quote")
		data, err := getQuoteHandler(w, r)
		if err != nil {
			log.Println("failure - processing the request: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonQuote, err := json.Marshal(data.USDBRL.Bid)
		if err != nil {
			log.Println("failure - marshaling json: " + err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
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

func getQuoteHandler(w http.ResponseWriter, r *http.Request) (*Quote, error) {
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

func getQuote(ctx context.Context, url string) (*Quote, error) {
	var quote *Quote
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
