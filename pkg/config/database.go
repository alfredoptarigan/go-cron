package config

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Database struct {
	*gorm.DB
}

type DatabasePrimadax Database
type DatabaseManApi Database
type Mt5Database Database
type DemoMt5Database Database

func InitDatabase() *gorm.DB {

	dbHost := DbHost
	dbName := DbName
	dbUser := DbUser
	dbPass := DbPass
	dbPort := DbPort

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
