package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"time"
)

func LogMiddleware() func(ctx *fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		f, _ := os.OpenFile(fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02")), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		defer f.Close()

		newLog := logrus.New()
		newLog.SetLevel(logrus.InfoLevel)
		newLog.SetFormatter(&logrus.JSONFormatter{})
		if os.Getenv("APP_ENV") == "production" {
			newLog.SetOutput(f)
		} else {
			newLog.SetOutput(io.MultiWriter(f, os.Stdout))
		}
		newLog.WithFields(logrus.Fields{
			"method": ctx.Method(),
			"uri":    ctx.OriginalURL(),
			"ip":     ctx.IP(),
			"code":   ctx.Response().StatusCode(),
		}).Log(logrus.InfoLevel, "Request")

		err := ctx.Next()
		if err != nil {
			return err
		}
		return nil
	}
}
