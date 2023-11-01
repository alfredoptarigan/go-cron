package config

import "gocron.com/m/pkg/helper"

var (
	DbHost    = helper.SetConfig("DB_GIC_API_HOST", "localhost")
	DbName    = helper.SetConfig("DB_GIC_API_NAME", "")
	DbUser    = helper.SetConfig("DB_GIC_API_USER", "")
	DbPass    = helper.SetConfig("DB_GIC_API_PASS", "")
	DbPort    = helper.SetConfig("DB_GIC_API_PORT", "3306")
	AppEnv    = helper.SetConfig("APP_ENV", "development")
	AppPort   = helper.SetConfig("APP_PORT", "9090")
	JwtSecret = helper.SetConfig("JWT_SECRET", "")
	JwtExpire = helper.SetConfig("JWT_EXPIRED", "")
)
