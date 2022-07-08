package database

import (
	"strconv"
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

	user.Token, err = middlewares.CreateToken(user.ID)
	if err != nil {
		return nil, err
	}

	if err := DB.Save(user).Error; err != nil {
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

func GetUser(user *models.User) (interface{}, error) {
	var err error
	if err = DB.First(user, user.ID).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func CreateUser(user *models.User) (interface{}, error) {
	strUser_role_id := strconv.FormatUint(uint64(user.User_role_id), 10)

	sql := `INSERT INTO users(user_role_id, username, name, phone, email, address, password, created_at, updated_at)
			VALUES(` + strUser_role_id + `, "` + user.Username + `", "` + user.Name + `", "` + user.Phone + `",
			"` + user.Email + `", "` + user.Address + `", "` + user.Password + `", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	if err := DB.Exec(sql).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func DeleteUser(user *models.User) (interface{}, error) {
	var err error
	if err = DB.Where("id = ?", user.ID).Delete(user).Error; err != nil {
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
