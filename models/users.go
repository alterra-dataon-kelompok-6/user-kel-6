package models

type User_role struct {
	ID   uint   `json:"id" query:"id" gorm:"primarykey;autoIncrement"`
	Role string `gorm:"type:varchar(200)" json:"role"`
}

type User struct {
	Common
	User_role_id uint      `json:"user_role_id" query:"user_role_id" form:"user_role_id"`
	Username     string    `gorm:"type:varchar(100);unique" query:"username" json:"username" form:"username"`
	Name         string    `gorm:"type:varchar(200);" query:"name" json:"name" form:"name"`
	Phone        string    `gorm:"type:varchar(20);" query:"phone" json:"phone" form:"phone"`
	Email        string    `gorm:"type:varchar(200);unique" query:"email" json:"email" form:"email"`
	Address      string    `gorm:"type:varchar(200);" query:"address" json:"address" form:"address"`
	Password     string    `gorm:"type:varchar(200);" query:"password" json:"password" form:"password"`
	Token        string    `gorm:"type:varchar(700);" query:"token" json:"token" form:"token"`
	User_role    User_role `gorm:"foreignKey:User_role_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
