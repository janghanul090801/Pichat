package handlers

import (
	"Pichat/api/middleware"
	"Pichat/api/presenter"
	"Pichat/pkg/auth"
	"Pichat/pkg/entities"
	"Pichat/pkg/users"
	"github.com/gofiber/fiber/v2"
)

func RegisterUser(userService users.Service, authService auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input, err := GetBody[presenter.CreateUserRequest](c)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		hashed, err := authService.HashPassword(input.Password)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		user := &entities.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: hashed,
		}

		user, err = userService.InsertUser(user)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		return c.JSON(user.ToResponse())
	}
}

func GetUserInfo(userService users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id, err := c.ParamsInt("id")
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		user, err := userService.GetUserByID(id)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		return c.JSON(user.ToResponse())
	}
}

func LoginUser(userService users.Service, authService auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		input, err := GetBody[presenter.LoginUserRequest](c)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		user, err := userService.GetUserByEmail(input.Email)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		compare := authService.CheckPassword(input.Password, user.Password)
		if !compare {
			return fiber.NewError(401, "Invalid credentials")
		}

		jwtToken, err := authService.GenerateJWToken(user.Email)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		return c.JSON(fiber.Map{"token": jwtToken})
	}
}

// private

func GetMyInfo(userService users.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := middleware.GetCurrentUserEmail(c)

		user, err := userService.GetUserByEmail(email)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		return c.JSON(user.ToResponse())
	}
}

func UpdateUser(userService users.Service, authService auth.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		email := middleware.GetCurrentUserEmail(c)
		user, err := userService.GetUserByEmail(email)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		input, err := GetBody[presenter.UpdateUserRequest](c)
		if err != nil {
			return fiber.NewError(400, err.Error())
		}

		hashed, err := authService.HashPassword(input.Password)

		updateUser := &entities.User{
			Name:     input.Name,
			Email:    input.Email,
			Password: hashed,
		}

		user, err = userService.UpdateUser(updateUser, user.ID)
		if err != nil {
			return fiber.NewError(500, err.Error())
		}

		return c.JSON(user.ToResponse())
	}
}
