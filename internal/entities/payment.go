package entities

const (
	Credit = "credit"
	Paypal = "paypal"
	Cash   = "cash"
)

type Payment interface {
	GetType() string
	GetMethodName() string
	GetToken() string
}
