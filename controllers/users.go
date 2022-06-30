package controllers

import (
	"net/http"
	"strconv"
	"time"

	"user-kel-6/config"
	"user-kel-6/lib/database"
	"user-kel-6/models"

	"github.com/labstack/echo/v4"
)

var users []models.User
var guna = models.User{}
var DB = config.DB

func LoginUsersController(c echo.Context) error {
	c.Bind(&guna)

	guna.Email = c.FormValue("email")
	guna.Password = c.FormValue("password")
	users, e := database.LoginUsers(&guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success login",
		"users":  users,
	})
}

// get all users
func GetUsersController(c echo.Context) error {
	users, e := database.GetUser(&guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	ID, _ := strconv.Atoi(c.Param("ID"))
	guna.ID = int(ID)

	if err := DB.First(&guna, ID).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"users":    &guna,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	usr := new(models.User)
	if err := c.Bind(usr); err != nil {
		return err
	}

	usr.User_role_id, _ = strconv.Atoi(c.FormValue("user_role_id"))
	usr.Username = c.FormValue("username")
	usr.Name = c.FormValue("name")
	usr.Phone = c.FormValue("phone")
	usr.Email = c.FormValue("email")
	usr.Address = c.FormValue("address")
	usr.Password = c.FormValue("password")

	strUser_role_id := strconv.Itoa(int(usr.User_role_id))

	sql := `INSERT INTO users(user_role_id, username, name, phone, email, address, password, created_at, updated_at)
			VALUES(` + strUser_role_id + `, "` + usr.Username + `", "` + usr.Name + `", "` + usr.Phone + `",
			"` + usr.Email + `", "` + usr.Address + `", "` + usr.Password + `", CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)`

	if err := DB.Exec(sql).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create an user",
		"users":    usr,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	reqId, _ := strconv.Atoi(c.Param("ID"))

	if err := DB.Where("ID = ?", reqId).Delete(&guna).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "success delete ID "+strconv.Itoa(reqId))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return err
	}

	ID, _ := strconv.Atoi(c.Param("ID"))

	user.ID = ID
	user.Username = c.FormValue("username")
	user.Name = c.FormValue("name")
	user.Phone = c.FormValue("phone")
	user.Email = c.FormValue("email")
	user.Address = c.FormValue("address")
	user.Password = c.FormValue("password")
	user.UpdatedAt = time.Now()

	doUpdate := DB.Table("users").Where("ID IN ?", []int{ID}).
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
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an user",
		"user":     user,
	})
}
