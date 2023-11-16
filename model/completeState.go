package model

import (
	"fmt"
)

type completeState struct{
	purchaseProcess *PurchaseProcess
}
func(p *orderState) Complete(){
	fmt.Println("Delivery Completed")
}
