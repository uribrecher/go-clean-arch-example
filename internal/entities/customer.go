package entities

type Customer interface {
	GetFullName() string
	GetPhoneNumber() string
	GetShippingAddress() ShippingDetails
	GetEmail() string
	GetPayment() Payment
}
