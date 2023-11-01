package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	fiberCompress "github.com/gofiber/fiber/v2/middleware/compress"
	fiberCors "github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
	"gocron.com/m/pkg/config"
	"gocron.com/m/pkg/cron"
	"gocron.com/m/pkg/middleware"
	"gocron.com/m/pkg/models"
	"gocron.com/m/pkg/routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func main() {
	utcLocation := time.UTC
	time.Local = utcLocation

	fiberConfig := config.NewFiberConfig()
	recoverMiddleware := middleware.Recover()
	LogMiddleware := middleware.LogMiddleware()

	app := fiber.New(fiberConfig)

	db := openDB()
	defer func() {
		sqlDB, err := db.DB()
		if err != nil {
			log.Fatalf("Error while connecting to database %s", err)
		}

		sqlDB.Close()
	}()

	//Migrate database
	err := db.AutoMigrate(&models.User{}, &models.UserLog{})
	if err != nil {
		logrus.Fatalf("Error while migrating database %s", err)
	}

	//Jika tidak di production/staging maka tidak perlu recover panic, karena akan memperlihatkan error yang sebenarnya
	if config.AppEnv == "production" {
		app.Use(recoverMiddleware) //recover panic
	}

	if config.AppEnv == "staging" {
		app.Use(fiberCors.New(fiberCors.Config{
			AllowOrigins: "*",
			AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
			AllowHeaders: "Origin, Content-Type, Accept, Authorization, Lang, lang, Accept-Encoding",
		}))
		app.Use(fiberCompress.New(fiberCompress.Config{
			Level: fiberCompress.LevelBestSpeed,
		}))
		app.Use(recoverMiddleware)
	}

	cron.StartCron(db)

	app.Use(LogMiddleware)
	routes.InitializeRouteV1(app)

	port := fmt.Sprintf(":%s", config.AppPort)
	err = app.Listen(port)
	if err != nil {
		logrus.Fatalf("Error while starting server %s", err)
	}

}

func openDB() *gorm.DB {

	dbHost := config.DbHost
	dbName := config.DbName
	dbUser := config.DbUser
	dbPass := config.DbPass
	dbPort := config.DbPort

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
	}

	//config database
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Error while connecting to database %s", err)
	}

	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetMaxOpenConns(20)
	sqlDB.SetConnMaxLifetime(60 * time.Minute)
	sqlDB.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
