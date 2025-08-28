package handlers

import (
	"Pichat/api/validator"
	"github.com/gofiber/fiber/v2"
)

func GetBody[T any](c *fiber.Ctx) (T, error) {
	var input T
	if err := c.BodyParser(&input); err != nil {
		return input, fiber.NewError(400, err.Error())
	}
	err := validator.Validate.Struct(input)
	if err != nil {
		return input, fiber.NewError(400, err.Error())
	}

	return input, nil
}
