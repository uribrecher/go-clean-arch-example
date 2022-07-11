package main

import (
	"clean/internal/application/v3"
	"clean/internal/wrappers/recordstore"
)

func main() {
	svc := recordstore.New()
	app := v3.NewApp(svc)
	app.PlaceOrder()
}
