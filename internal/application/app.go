package application

import "clean/thirdparty/recordstore"

type App struct {
	svc recordstore.Service
}

func (a App) PlaceOrder() {
	order := a.svc.CreateRecord()
	a.svc.StoreRecord(order)
}

func NewApp() App {
	return App{svc: recordstore.New()}
}
