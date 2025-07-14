package main

import (
	"AUTH_IN_GO/app"
	config "AUTH_IN_GO/config/env"
)

func main() {

	config.Load()

	cfg := app.NewConfig() // Set the server to listen on port 8080
	app := app.NewApplication(cfg)

	app.Run()
}