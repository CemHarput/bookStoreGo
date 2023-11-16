package model

import (
	"github.com/CemHarput/bookStoreGo/interfaces"
	"gorm.io/gorm"
)
type PurchaseProcess struct {
	gorm.Model
	defaultState interfaces.PurchaseState
	orderState interfaces.PurchaseState
	cargoState interfaces.PurchaseState
	deliveryState interfaces.PurchaseState
	completeState interfaces.PurchaseState
	PurchaseOrder  *PurchaseOrder `gorm:"foreignKey:PurchaseProcessID"`
	PurchaseOrderID uint
}
func (purchaseState *PurchaseProcess) setState(state interfaces.PurchaseState) {
	purchaseState.defaultState = state
}
func (purchaseState *PurchaseProcess) Order(){
	purchaseState.defaultState.Order()
}
func (purchaseState *PurchaseProcess) Refund(){
	purchaseState.defaultState.Refund()
}
