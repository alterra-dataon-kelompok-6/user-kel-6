package controllers

import (
	"net/http"
	"strconv"
	"time"

	"user-kel-6/config"
	"user-kel-6/lib/database"
	"user-kel-6/middlewares"
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
	user, e := database.LoginUsers(guna)

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"email":    user.Email,
		"password": user.Password,
		"token":    user.Token,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	c.Bind(&guna)

	ID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	guna.ID = uint(ID)

	use, e := database.GetUser(&guna)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	if use == false {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "user with requested ID was not found",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"ID":       guna.ID,
		"username": guna.Username,
		"name":     guna.Name,
		"phone":    guna.Phone,
		"email":    guna.Email,
		"address":  guna.Address,
		"password": guna.Password,
	})
}

// create new user
func CreateUserController(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	u.User_role_id = 2
	u.Username = c.FormValue("username")
	u.Name = c.FormValue("name")
	u.Phone = c.FormValue("phone")
	u.Email = c.FormValue("email")
	u.Address = c.FormValue("address")
	u.Password = c.FormValue("password")

	// Validasi input harus terisi semuanya
	validateInput, field := database.ValidateInput(u)
	if !validateInput {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: field + " field is required !!",
		})
	}

	// validasi registered email
	validateEmail, err := database.ValidateEmail(u.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if !validateEmail {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Email is already registered, please use another email",
		})
	}

	// validasi registered phone number
	validatePhone, err := database.ValidatePhone(u.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validatePhone {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Phone number is already registered, please use another phone number",
		})
	}

	// validasi registered phone number
	validateUser, err := database.ValidateUser(u.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validateUser {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Username is already registered, please use another phone number",
		})
	}

	// input ke database
	_, e := database.CreateUser(u)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"MESSAGE":  "success create an user",
		"username": u.Username,
		"name":     u.Name,
		"phone":    u.Phone,
		"email":    u.Email,
		"address":  u.Address,
		"password": u.Password,
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	c.Bind(&guna)

	guna.ID = middlewares.GetUserIdFromToken(c)
	guna.Username = c.FormValue("username")
	guna.Name = c.FormValue("name")
	guna.Phone = c.FormValue("phone")
	guna.Email = c.FormValue("email")
	guna.Address = c.FormValue("address")
	guna.Password = c.FormValue("password")
	guna.UpdatedAt = time.Now()

	// Validasi input harus terisi semuanya
	validateInput, field := database.ValidateInput(&guna)
	if !validateInput {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: field + " field is required !!",
		})
	}

	// validasi registered email
	validateEmail, err := database.ValidateEmail(guna.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	} else if !validateEmail {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Email is already registered, please use another email",
		})
	}

	// validasi registered phone number
	validatePhone, err := database.ValidatePhone(guna.Phone)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validatePhone {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Phone number is already registered, please use another phone number",
		})
	}

	// validasi registered phone number
	validateUser, err := database.ValidateUser(guna.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	} else if !validateUser {
		return c.JSON(http.StatusBadRequest, models.Response{
			Status:  "fail",
			Message: "Username is already registered, please use another phone number",
		})
	}

	// input ke database
	_, e := database.CreateUser(&guna)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"MESSAGE":    "success update an user",
		"ID":         guna.ID,
		"username":   guna.Username,
		"name":       guna.Name,
		"phone":      guna.Phone,
		"email":      guna.Email,
		"address":    guna.Address,
		"password":   guna.Password,
		"updated at": guna.UpdatedAt,
	})
}

// FUTURE DEVELOPMENT - FOR ADMINS ONLY
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
