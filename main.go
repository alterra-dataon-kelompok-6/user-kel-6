package main

import (
	"user-kel-6/routes"

	m "user-kel-6/middlewares"
)

func main() {
	// create a new echo instance
	e := routes.New()

	// implement middleware logger
	m.LogMiddlewares(e)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8088"))
}
