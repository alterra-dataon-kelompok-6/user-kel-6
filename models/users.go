package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int    `gorm:"primaryKey"`
	User_role_id int    `json:"user_role_id" form:"user_role_id"`
	Username     string `gorm:"type:varchar(100);" json:"username" form:"username"`
	Name         string `gorm:"type:varchar(200);" json:"name" form:"name"`
	Phone        string `gorm:"type:varchar(20);" json:"phone" form:"phone"`
	Email        string `gorm:"type:varchar(200);" json:"email" form:"email"`
	Address      string `gorm:"type:varchar(200);" json:"address" form:"address"`
	Password     string `gorm:"type:varchar(200);" json:"password" form:"password"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Token        string         `gorm:"type:varchar(700);" json:"token" form:"token"`
}
