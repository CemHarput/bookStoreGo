package main

import "github.com/shopspring/decimal"

// Define the Book struct
type Book struct {
    CommonFields
    Title  string `json:"title"`
    Author string `json:"author"`
	Price  decimal.Decimal `json:"price"`
	StockQuantity int `json:"stockQuantity"`
}