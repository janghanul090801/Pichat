package routes

import (
	"Pichat/api/handlers"
	"Pichat/api/middleware"
	"Pichat/pkg/auth"
	"Pichat/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func UserRouter(app fiber.Router, userService users.Service, authService auth.Service) {
	app = app.Group("/users")
	app.Post("/", handlers.RegisterUser(userService, authService))
	app.Get("/:id", handlers.GetUserInfo(userService))
	app.Get("/login", handlers.LoginUser(userService, authService))

	private := app.Group("/private", middleware.JWTAuthMiddleware)
	private.Get("/", handlers.GetMyInfo(userService))
	private.Put("/", handlers.UpdateUser(userService, authService))
}
