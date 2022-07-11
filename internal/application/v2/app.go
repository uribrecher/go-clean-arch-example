package v2

// This file contains compilation error intentionally

import recordstore "clean/thirdparty/recordstore/v2"

type App struct {
	svc recordstore.Service
}

func (a App) PlaceOrder() {
	order := a.svc.CreateRecord()
	a.svc.StoreRecord(order)
}

func (a App) PlaceOrder2() {
	ordersTable := a.svc.GetTable("orders")
	ordersTable.StoreRecord(order)
}
