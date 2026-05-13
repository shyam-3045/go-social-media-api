package users

import (
	"net/http"
	dto "socio/internals/Dto"
	"socio/internals/validator"
	"socio/services/users"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
		"ok":   true,
		"data": user,
	})
}

func GetByID(c fiber.Ctx) error {
	ctx := c.Context()

	id := c.Params("id")

	userID, err := uuid.Parse(id)

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"ok":    false,
			"error": "incorrect user id ",
		})
	}

	user, err := users.GetById(ctx, userID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"ok":    false,
				"error": "user not found  ",
			})

		}

	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"ok":   true,
		"data": user,
	})

}
