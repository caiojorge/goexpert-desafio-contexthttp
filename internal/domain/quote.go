package domain

import (
	"log"
	"strconv"
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

func (q *Quote) GetBid() float64 {
	bid, err := strconv.ParseFloat(q.USDBRL.Bid, 64)
	if err != nil {
		log.Fatalf("Failed to convert string to float64: %v", err)
	}
	return bid
}
