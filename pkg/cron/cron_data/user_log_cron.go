package cron_data

import "gocron.com/m/pkg/models"

type UserLogCron interface {
	GetUserLog() ([]models.User, error)
	StoreUserLogRegister(data models.User) (models.UserLog, error)
}
