package model

import (
	"fmt"
)

type orderState struct{
	purchaseProcess *PurchaseProcess
}
func(p *orderState) StartOrder(){
	fmt.Println("Order process is started")
	p.purchaseProcess.setState(p.purchaseProcess.cargoState)
}
func(p *orderState) CancelOrder(){
	fmt.Println("Order process is cancelled")
	p.purchaseProcess.setState(p.purchaseProcess.defaultState)
} 