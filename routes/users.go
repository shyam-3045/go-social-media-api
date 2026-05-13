package routes

import (
	"github.com/gofiber/fiber/v3"
	"socio/controllers/users"
)

func Users(r fiber.Router) {
	usersRoutes := r.Group("/users")

		usersRoutes.Post("/",users.Add)
		// users.Get("/",nil)

		// users.Get("/:id",nil)
		// users.Delete("/:id",nil)

}