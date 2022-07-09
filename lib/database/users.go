package database

import (
	"errors"
	"user-kel-6/config"
	"user-kel-6/middlewares"
	"user-kel-6/models"
)

var DB = config.DB

func LoginUsers(user models.User) (*models.User, error) {
	var err error

	if err := DB.Where("email = ? AND password = ?", user.Email, user.Password).Find(&user).Error; err != nil {
		return nil, err
	}

	if user.ID == 0 {
		return nil, errors.New("record not found")
	}

	user.Token, err = middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	doUpdate := DB.Debug().Table("users").Where("id = ?", user.ID).
		Updates(map[string]interface{}{
			"token": user.Token,
		})

	if err := doUpdate.Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(user *models.User) (interface{}, error) {
	var err error
	if err = DB.First(user, user.ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	if err := DB.Save(&user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUser(user *models.User) (interface{}, error) {
	doUpdate := DB.Table("users").Where("id IN ?", []uint{user.ID}).
		Updates(map[string]interface{}{
			"username":   user.Username,
			"name":       user.Name,
			"phone":      user.Phone,
			"email":      user.Email,
			"address":    user.Address,
			"password":   user.Password,
			"updated_at": user.UpdatedAt,
		})

	if err := doUpdate.Error; err != nil {
		return nil, err
	}

	return user, nil
}

func ValidateEmail(email string) (bool, error) {
	var user models.User
	query := config.DB.Where("email = ?", email).Find(&user)
	if query.Error != nil {
		return false, query.Error
	} else if query.RowsAffected != 0 {
		return false, nil
	}
	return true, nil
}

func ValidatePhone(phone string) (bool, error) {
	var user models.User
	query := config.DB.Where("phone = ?", phone).Find(&user)
	if query.Error != nil {
		return false, query.Error
	} else if query.RowsAffected != 0 {
		return false, nil
	}
	return true, nil
}

func ValidateUser(username string) (bool, error) {
	var user models.User
	query := config.DB.Where("username = ?", username).Find(&user)
	if query.Error != nil {
		return false, query.Error
	} else if query.RowsAffected != 0 {
		return false, nil
	}
	return true, nil
}

func ValidateInput(user *models.User) (bool, string) {
	switch {
	case user.Name == "":
		return false, "name"
	case user.Email == "":
		return false, "email"
	case user.Phone == "":
		return false, "phone"
	case user.Password == "":
		return false, "password"
	case user.Address == "":
		return false, "address"
	}

	return true, ""
}

// FUTURE DEVELOPMENT - FOR ADMINS ONLY
func DeleteUser(user *models.User) (interface{}, error) {
	var err error
	if err = DB.Where("id = ?", user.ID).Delete(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func GetUsers(user *[]models.User) (interface{}, error) {
	var err error
	if err = DB.Find(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
