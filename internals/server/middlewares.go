package server

import (
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func middlewares(app *fiber.App){
	app.Use(logger.New())
}

