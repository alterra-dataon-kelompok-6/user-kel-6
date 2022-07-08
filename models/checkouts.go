package models

import (
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type Checkout_statu struct {
	ID        uint   `gorm:"primaryKey"`
	Status    string `json:"status" form:"status" gorm:"type:varchar(100);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Checkout struct {
	ID                 uint            `gorm:"primaryKey"`
	Cart_id            uint            `json:"cart_id" form:"cart_id"`
	Total_qty          int             `json:"total_qty" form:"total_qty"`
	Total_price        decimal.Decimal `json:"total_price" form:"total_price" gorm:"type:decimal(25,2);"`
	Checkout_status_id uint            `json:"checkout_status_id" form:"checkout_status_id"`
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt `gorm:"index"`
}
