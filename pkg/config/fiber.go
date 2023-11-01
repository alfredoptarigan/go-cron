package config

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	fiberCompress "github.com/gofiber/fiber/v2/middleware/compress"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	"gocron.com/m/pkg/exception"
)

func NewFiberConfig() fiber.Config {
	return fiber.Config{
		ErrorHandler:      exception.ErrorHandler,
		EnablePrintRoutes: true,
		JSONEncoder:       json.Marshal,
		JSONDecoder:       json.Unmarshal,
	}
}

func NewFiberCorsConfig() fiberCors.Config {
	return fiberCors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization, Lang, lang, Accept-Encoding",
	}
}

func NewFiberCompressionConfig() fiberCompress.Config {
	return fiberCompress.Config{
		Level: fiberCompress.LevelBestSpeed,
	}
}
