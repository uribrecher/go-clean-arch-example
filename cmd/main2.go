package main

import (
	"clean/internal/application/v3"
	recordstore "clean/internal/wrappers/recordstore/v2"
)

func main() {
	svc := recordstore.New()
	app := v3.NewApp(svc)
	app.PlaceOrder()
}
