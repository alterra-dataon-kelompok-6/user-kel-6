package config

import (
	"fmt"
	"user-kel-6/constants"
	"user-kel-6/models"

	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	InitDB()
	InitMigrate()
	CreateUserRole()
}

func InitDB() {
	var err error
	DB, err = gorm.Open(mysql.Open(ConnString()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func ConnString() string {
	config := models.Config{
		DB_Username: constants.DB_USERNAME,
		DB_Password: constants.DB_PASSWORD,
		DB_Port:     constants.DB_PORT,
		DB_Host:     constants.DB_HOST,
		DB_Name:     constants.DB_NAME,
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	return connectionString
}

func CreateUserRole() {
	DB.Exec(`INSERT IGNORE INTO user_roles(id, role) VALUES (1, 'administrator');`)
	DB.Exec(`INSERT IGNORE INTO user_roles(id, role) VALUES (2, 'customer');`)
}

func InitMigrate() {
	DB.AutoMigrate(&models.User_role{})
	DB.AutoMigrate(&models.User{})
}
