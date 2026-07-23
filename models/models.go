package models

type Stock struct {
	StockID     string  `json:"stockid"`
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Company int     `json:"company"`
}