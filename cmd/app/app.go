package app

import (
	"log"
	"socio/internals/database"
	"socio/internals/server"
)

func Setup() {
	database.Config()

	app := server.New()

	if err := app.Listen(":8080");err != nil {
		log.Fatalf("Error starting server , error : %v" ,err)
	}
}