package server

import "github.com/gofiber/fiber/v3"

func errorHandler(c fiber.Ctx ,e error) error{

	msg:=e.Error()

	return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
		"ok":false,
		"error":msg,
	})
}

var notFoundHandler = func(c fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
		"ok":false,
		"error":"resource not found",
	})
} 