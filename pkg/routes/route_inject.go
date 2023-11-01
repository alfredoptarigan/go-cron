package routes

import (
	"github.com/gofiber/fiber/v2"
	"gocron.com/m/pkg/injection"
)

func InitializeRouteV1(app *fiber.App) {
	api := app.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			userController := injection.InitializeUserController()
			userController.Route(v1, app)
		}
	}
}
