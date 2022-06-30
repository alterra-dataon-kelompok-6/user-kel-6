package database

import (
	"user-kel-6/config"
	"user-kel-6/middlewares"
	"user-kel-6/models"
)

var DB = config.DB

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUser(user *models.User) (interface{}, error) {
	var err error
	if err = DB.Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
