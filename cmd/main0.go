package main

import "clean/internal/application"

func main() {
	app := application.NewApp()
	app.PlaceOrder()
}
