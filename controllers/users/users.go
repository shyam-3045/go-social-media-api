package users

import (
	dto "socio/internals/Dto"
	"socio/internals/validator"
	"socio/services/users"

	"github.com/gofiber/fiber/v3"
)

func Add(c fiber.Ctx) error {

	ctx := c.Context()

	var body dto.UserCreate

	if err := c.Bind().Body(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"invalid request body",
		)
	}

	if err := validator.Users(body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			err.Error(),
		)
	}

	user, err := users.Create(ctx, body)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ok": true,
		"data": user,
	})
}