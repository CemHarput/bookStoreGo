package model

import (
	"gorm.io/gorm"
)
type PurchaseOrderBook struct {
    gorm.Model
    PurchaseOrderID uint
    BookID          uint
}