package routes

import (
	"socio/controllers/friendships"

	"github.com/gofiber/fiber/v3"
)

func FriendShips(r fiber.Router) {
	friendship := r.Group("/friendship")

	friendship.Post("/",friendships.Create)


}