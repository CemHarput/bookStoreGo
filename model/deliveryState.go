package model

import (
	"fmt"
)

type deliveryState struct{
	purchaseProcess *PurchaseProcess
}
func(p *orderState) StartDelivery(){
	fmt.Println("Delivery process is started")
	p.purchaseProcess.setState(p.purchaseProcess.completeState)
}
func(p *orderState) CancelDelivery(){
	fmt.Println("Delivery process is cancelled")
	p.purchaseProcess.setState(p.purchaseProcess.defaultState)
}