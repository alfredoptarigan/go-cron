package models

import (
	"gorm.io/gorm"
	"time"
)

type UserLog struct {
	Uuid        string         `json:"uuid" gorm:"primaryKey, type:uuid;"`
	UserUuid    string         `json:"user_uuid"`
	IpAddress   string         `json:"ip_address" gorm:"type:varchar(255)"`
	Description string         `json:"description" gorm:"type:varchar(255)"`
	CreatedAt   time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt   time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

func (u *UserLog) TableName() string {
	return "user_logs"
}
