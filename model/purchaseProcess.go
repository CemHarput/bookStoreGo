package model

import (
	"github.com/CemHarput/bookStoreGo/interfaces"
)
type PurchaseProcess struct {
	defaultState interfaces.PurchaseState
	orderState interfaces.PurchaseState
	cargoState interfaces.PurchaseState
	deliveryState interfaces.PurchaseState
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
