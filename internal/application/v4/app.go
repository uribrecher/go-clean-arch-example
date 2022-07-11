package v4

import (
	"clean/internal/ports/customers"
	"clean/internal/ports/orders"
	"clean/internal/ports/products"
	"clean/internal/ports/shipping"
)

type App struct {
	ordersPort     orders.Service
	customersPort  customers.Service
	productCatalog products.Service
	shippingPort   shipping.Service
}

func (a App) PlaceOrder(
	customerID string,
	productsIDs []string,
) {
	customer := a.customersPort.GetCustomer(customerID)
	order := a.ordersPort.NewOrder()
	order.SetShippingAddress(customer.GetShippingAddress())

	for _, id := range productsIDs {
		product := a.productCatalog.GetProductByID(id)
		order.AddProduct(product)
	}
	order.AddPayment(customer.GetPayment())
	a.ordersPort.SaveOrder(order)
}
