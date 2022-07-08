package models

type User_role struct {
	ID   uint   `json:"id" gorm:"primarykey;autoIncrement"`
	Role string `gorm:"type:varchar(200)" json:"role"`
}

type User struct {
	ID           uint      `json:"id" gorm:"primarykey;autoIncrement"`
	User_role_id uint      `json:"user_role_id" form:"user_role_id"`
	Username     string    `gorm:"type:varchar(100);" json:"username" form:"username"`
	Name         string    `gorm:"type:varchar(200);" json:"name" form:"name"`
	Phone        string    `gorm:"type:varchar(20);" json:"phone" form:"phone"`
	Email        string    `gorm:"type:varchar(200);unique" json:"email" form:"email"`
	Address      string    `gorm:"type:varchar(200);" json:"address" form:"address"`
	Password     string    `gorm:"type:varchar(200);" json:"password" form:"password"`
	Token        string    `gorm:"type:varchar(700);" json:"token" form:"token"`
	User_role    User_role `gorm:"foreignKey:User_role_id; constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Times
}
