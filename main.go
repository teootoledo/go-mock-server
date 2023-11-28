package main

import (
	"mock-server/internal/app"
)

//	@title			Mock Server API
//	@version		1.0.0
//	@description	API for mocking responses

//	@contact.name	Teo Martin Toledo
//	@contact.email	teootoledo@gmail.com

//	@host		localhost:8080
//	@BasePath	/v1

func main() {
	app.NewApp().
		Setup().
		InitServer()
}
