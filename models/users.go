package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           int    `gorm:"primaryKey"`
	User_role_id int    `json:"user_role_id" form:"user_role_id"`
	Username     string `json:"username" form:"username"`
	Name         string `json:"name" form:"name"`
	Phone        string `json:"phone" form:"phone"`
	Email        string `json:"email" form:"email"`
	Address      string `json:"address" form:"address"`
	Password     string `json:"password" form:"password"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
	Token        string         `json:"token" form:"token"`
}
