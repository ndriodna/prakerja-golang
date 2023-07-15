package routes

import (
	"day_eight/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	users := e.Group("/users")
	users.GET("", controllers.UserIndex)
	users.GET("/:id", controllers.UserShow)
	users.POST("/create", controllers.UserCreate)
	users.PATCH("/:id/update", controllers.UserUpdate)
	users.DELETE("/:id/delete", controllers.UserDelete)
	return e
}
