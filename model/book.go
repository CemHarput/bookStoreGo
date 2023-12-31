package model

import (
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

// Define the Book struct
type Book struct {
    gorm.Model
    CommonFields
    Title  string `json:"title"`
    Author string `json:"author"`
	Price  decimal.Decimal `json:"price"`
	StockQuantity int `json:"stockQuantity"`
	PurchaseOrders []*PurchaseOrder `gorm:"many2many:purchase_order_books;"`
}