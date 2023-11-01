package user

import (
	"github.com/gofiber/fiber/v2"
	"gocron.com/m/pkg/dto/user"
	"gocron.com/m/pkg/helper"
	userService "gocron.com/m/pkg/services/user"
)

type userControllerImpl struct {
	userService userService.UserService
}

func (u *userControllerImpl) Route(router fiber.Router, app *fiber.App) {
	router.Post("/register", u.RegisterUser)
}

func (u *userControllerImpl) RegisterUser(ctx *fiber.Ctx) error {
	var userDTO user.RegisterUserDTO

	if err := ctx.BodyParser(&userDTO); err != nil {
		return ctx.Status(400).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	if err := u.userService.RegisterUser(userDTO); err != nil {
		return ctx.Status(400).JSON(helper.ApiResponseError(err.Error()))
	}

	return ctx.JSON(helper.ApiResponseSuccess("success", nil))
}

func NewUserController(userService userService.UserService) UserController {
	return &userControllerImpl{
		userService: userService,
	}
}
