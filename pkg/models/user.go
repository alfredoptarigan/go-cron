package models

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type User struct {
	Uuid            string         `json:"uuid" gorm:"primaryKey, type:uuid;"`
	Name            string         `json:"name" gorm:"type:varchar(255)"`
	Email           string         `json:"email" gorm:"type:varchar(255)"`
	EmailVerifiedAt sql.NullTime   `json:"email_verified_at" gorm:"type:timestamp"`
	Password        string         `json:"password" gorm:"type:varchar(255)"`
	CreatedAt       time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt       time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt       gorm.DeletedAt `json:"deleted_at"`
	IsCron          bool           `json:"is_cron" gorm:"default:false"`
}

func (u *User) TableName() string {
	return "users"
}
