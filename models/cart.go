package models

type Cart struct {
	Common
	CustomerID uint `json:"customer_id" gorm:"uniqueIndex"`
}
