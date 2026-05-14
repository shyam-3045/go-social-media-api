package routes

import (
	"github.com/gofiber/fiber/v3"
	"socio/controllers/users"
)

func Users(r fiber.Router) {
	usersRoutes := r.Group("/users")

		usersRoutes.Post("/",users.Add)
		usersRoutes.Get("/",users.Get)

		usersRoutes.Get("/:id",users.GetByID)
		usersRoutes.Delete("/:id",users.Delete)

}