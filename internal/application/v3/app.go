package v3

import (
	"clean/internal/ports/orders"
)

type App struct {
	svc orders.Service
}

func (a App) PlaceOrder() {
	order := a.svc.NewOrder()
	a.svc.SaveOrder(order)
}

func NewApp(s orders.Service) App {
	return App{svc: s}
}
