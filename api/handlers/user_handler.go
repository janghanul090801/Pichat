package handlers

import (
	"Pichat/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func HelloWorld(service user.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.JSON("Hello World")
	}
}
