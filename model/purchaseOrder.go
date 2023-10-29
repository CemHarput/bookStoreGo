package main

import (
	"time"

	"github.com/shopspring/decimal"
)
type PurchaseOrder struct {
    CommonFields
	Price  decimal.Decimal `json:"totalPrice"`
	StockQuantity int `json:"stockQuantity"`
	OrderDate time.Time `json:"orderDate"`
	Books       []Book `json:"books"`
}