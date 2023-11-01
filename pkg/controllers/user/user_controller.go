package user

import "github.com/gofiber/fiber/v2"

type UserController interface {
	Route(router fiber.Router, app *fiber.App)
	RegisterUser(ctx *fiber.Ctx) error
}
