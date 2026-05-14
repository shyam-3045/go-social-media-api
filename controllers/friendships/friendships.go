package friendships

import (
	dto "socio/internals/Dto"
	"socio/internals/validator"
	"socio/services/friendships"

	"github.com/gofiber/fiber/v3"
)

func Create(c fiber.Ctx) error {
	ctx := c.Context()

	var body dto.Friends

	if err := c.Bind().Body(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"input is not valid",
		)
	}

	if 	body.UserID == body.FriendID {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"user cannot friend themselves",
		)
	}

	if err := validator.Friends(&body); err != nil {
		return fiber.NewError(
			fiber.StatusBadRequest,
			"input is invalid",
		)
	}

	frnd, err := friendships.Add(ctx, body)

	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"ok":   true,
		"data": frnd})

}
