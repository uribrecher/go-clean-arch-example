package v2

import (
	"clean/internal/entities"
	"clean/thirdparty/recordstore"
)

type wrapper struct {
	impl recordstore.Service
}

type orderImpl struct{}

func (o orderImpl) GetID() string                                       { return "" }
func (o orderImpl) SetShippingAddress(details entities.ShippingDetails) {}
func (o orderImpl) AddProduct(product entities.Product)                 {}
func (o orderImpl) AddPayment(payment entities.Payment)                 {}

func (a wrapper) NewOrder() entities.Order {
	orderRecord := a.impl.CreateRecord()
	return toOrder(orderRecord)
}

func (a wrapper) SaveOrder(order entities.Order) {
	a.impl.StoreRecord(toRecord(order))
}

func toOrder(orderRecord interface{}) entities.Order {
	return orderImpl{}
}

func toRecord(order entities.Order) interface{} {
	return nil
}

func New() wrapper {
	return wrapper{impl: recordstore.New()}
}
