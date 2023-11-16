package model

import (
	"fmt"
)

type cargoState struct{
	purchaseProcess *PurchaseProcess
}
func(p *orderState) StartCargo(){
	fmt.Println("Cargo process is started")
	p.purchaseProcess.setState(p.purchaseProcess.deliveryState)
}
func(p *orderState) CancelCargo(){
	fmt.Println("Cargo process is cancelled")
	p.purchaseProcess.setState(p.purchaseProcess.defaultState)
}