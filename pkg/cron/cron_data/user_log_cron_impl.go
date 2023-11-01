package cron_data

import (
	"github.com/google/uuid"
	"gocron.com/m/pkg/models"
	"gorm.io/gorm"
	"time"
)

type userLogCronImpl struct {
	db *gorm.DB
}

func (u *userLogCronImpl) StoreUserLogRegister(data models.User) (models.UserLog, error) {
	userLog := models.UserLog{
		Uuid:        uuid.New().String(),
		UserUuid:    data.Uuid,
		IpAddress:   "1.1.1.1.1",
		Description: "User Has Been Register",
		CreatedAt:   time.Time{},
		UpdatedAt:   time.Time{},
	}

	err := u.db.Create(&userLog).Error
	if err != nil {
		return models.UserLog{}, err
	}

	// update is log cron on user
	err = u.db.Model(&data).Where("uuid = ?", data.Uuid).Update("is_cron", true).Error

	return userLog, nil
}

func (u *userLogCronImpl) GetUserLog() ([]models.User, error) {
	var userModel []models.User

	err := u.db.Debug().Find(&userModel).
		Where("created_at >= ?", time.Now().Add(-24*time.Hour)).
		Where("created_at <= ?", time.Now()).
		Where("email_verified_at IS NOT NULL").
		Where("is_cron = ? ", false).
		Find(&userModel).Error

	if err != nil {
		return []models.User{}, err
	}

	return userModel, nil
}

func NewUserLogCronImpl(db *gorm.DB) UserLogCron {
	return &userLogCronImpl{db: db}
}
