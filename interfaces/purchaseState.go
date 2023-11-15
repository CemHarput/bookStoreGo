package interfaces

type purchaseState interface {
	order()
	refund()
	checkThePurchaseProcess()
}