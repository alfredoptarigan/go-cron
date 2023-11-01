package middleware

import (
	"github.com/gofiber/fiber/v2"
	"gocron.com/m/pkg/helper"
)

func Recover() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		defer func() {
			if err := recover(); err != nil {
				helper.SaveErrorProduction(err.(error).Error(), err)
				err = ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
					"status":  "error",
					"message": "Something went wrong, please try again later",
				})
			}
		}()
		return ctx.Next()
	}
}
