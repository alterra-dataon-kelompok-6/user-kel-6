package routes

import (
	"user-kel-6/constants"
	c "user-kel-6/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	// create a new echo instance
	e := echo.New()

	// JWT Group
	r := e.Group("")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	// route users
	e.POST("/register", c.CreateUserController)
	e.POST("/login", c.LoginUsersController)

	// JWT: a user can access other users
	r.GET("/users/:id", c.GetUserController)

	// JWT: only edit him/hers
	r.PUT("/profile", c.UpdateUserController)

	return e
}
