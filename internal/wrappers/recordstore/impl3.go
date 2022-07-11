package recordstore

import (
	"clean/internal/entities"
	recordstore "clean/thirdparty/recordstore/v2"
)

type wrapper struct {
	impl recordstore.Service
}

type orderImpl struct{}

func (o orderImpl) GetID() string {return ""}
func (o orderImpl) SetShippingAddress(details entities.ShippingDetails) {}
func (o orderImpl) AddProduct(product entities.Product)                 {}
func (o orderImpl) AddPayment(payment entities.Payment)                 {}

func (a wrapper) NewOrder() entities.Order {
	return orderImpl{}
}

func (a wrapper) SaveOrder(order entities.Order) {
	ordersTable := a.impl.GetTable("orders")
	ordersTable.StoreRecord(toRecord(order))
}

func toRecord(order entities.Order) interface{} {
	return nil
}

func New() wrapper {
	return wrapper{impl: recordstore.New()}
}
