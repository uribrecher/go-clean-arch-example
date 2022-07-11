package entities

type Order interface {
	GetID() string
	SetShippingAddress(details ShippingDetails)
	AddProduct(product Product)
	AddPayment(payment Payment)
}
