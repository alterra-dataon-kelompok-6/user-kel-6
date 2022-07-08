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
	_, e := database.LoginUsers(&guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"email":    guna.Email,
		"password": guna.Password,
		"token":    guna.Token,
	})
}

// get all users
func GetUsersController(c echo.Context) error {
	sers, e := database.GetUsers(&users)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    sers,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	c.Bind(&guna)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	guna.ID = uint(ID)
	u, e := database.GetUser(&guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get an user",
		"users":    u,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	usr := new(models.User)
	if err := c.Bind(usr); err != nil {
		return err
	}

	int_role_id, _ := strconv.ParseUint(c.FormValue("user_role_id"), 10, 32)
	usr.User_role_id = uint(int_role_id)
	usr.Username = c.FormValue("username")
	usr.Name = c.FormValue("name")
	usr.Phone = c.FormValue("phone")
	usr.Email = c.FormValue("email")
	usr.Address = c.FormValue("address")
	usr.Password = c.FormValue("password")

	use, e := database.CreateUser(usr)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create an user",
		"users":    use,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	reqId, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	guna.ID = uint(reqId)
	_, e := database.GetUser(&guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, "success delete id "+strconv.FormatUint(reqId, 10))
}

// update user by id
func UpdateUserController(c echo.Context) error {
	c.Bind(&guna)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	guna.ID = uint(ID)
	guna.Username = c.FormValue("username")
	guna.Name = c.FormValue("name")
	guna.Phone = c.FormValue("phone")
	guna.Email = c.FormValue("email")
	guna.Address = c.FormValue("address")
	guna.Password = c.FormValue("password")
	guna.UpdatedAt = time.Now()

	use, e := database.UpdateUser(&guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success update an user",
		"user":     use,
	})
}
