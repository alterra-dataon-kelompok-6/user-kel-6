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
	e.POST("/login", c.LoginUsersController)

	r.GET("/users", c.GetUsersController)
	r.GET("/users/:ID", c.GetUserController)
	e.POST("/users", c.CreateUserController)
	r.DELETE("users/:ID", c.DeleteUserController)
	r.PUT("/users/:ID", c.UpdateUserController)

	return e
}
