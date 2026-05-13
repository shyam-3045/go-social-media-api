package app

import (
	"log"
	"socio/internals/database"
	"socio/internals/server"
	"socio/internals/validator"
)

func Setup() {
	database.Config()


	server.Setup()
	app := server.New()

	validator.Init()

	if err := app.Listen(":8080");err != nil {
		log.Fatalf("Error starting server , error : %v" ,err)
	}
}