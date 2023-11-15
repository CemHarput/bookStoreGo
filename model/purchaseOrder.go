package model

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)
type PurchaseOrder struct {
	gorm.Model
    CommonFields
	Price  decimal.Decimal `json:"totalPrice"`
	StockQuantity int `json:"stockQuantity"`
	OrderDate time.Time `json:"orderDate"`
	Books         []*Book         `gorm:"many2many:purchase_order_books;"`
	PurchaseProcess *PurchaseProcess `json:"PurchaseProcess"`
	 
}
func (purchaseOrder *PurchaseOrder) SetProcess(process *PurchaseProcess) {
	purchaseOrder.PurchaseProcess.setState(process) 
}

