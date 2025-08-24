package routes

import (
	"Pichat/api/handlers"
	"Pichat/pkg/user"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, service user.Service) {
	app = app.Group("/user")
	app.Get("/", handlers.HelloWorld(service))
}
