package models

import (
	"time"

	"gorm.io/gorm"
)

type Cart struct {
	ID          uint8 `gorm:"primaryKey"`
	Customer_id uint8 `json:"customer_id" form:"customer_id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Cart_item struct {
	ID         uint8 `gorm:"primaryKey"`
	Cart_id    uint8 `json:"cart_id" form:"cart_id"`
	Product_id uint8 `json:"product_id" form:"product_id"`
	Qty        int   `json:"qty" form:"qty"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
