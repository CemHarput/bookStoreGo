package interfaces

type PurchaseState interface {
	Order()
	Refund()
}