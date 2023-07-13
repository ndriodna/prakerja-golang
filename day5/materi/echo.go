package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id    int    `json:"id" form:"id"`
	Name  string `json:"name" form:"name"`
	Email string `json:"email" form:"email"`
}

func main() {
	// buat instance echo baru
	e := echo.New()
	user := User{1, "lele", "lele@mail.com"}
	// buat router
	e.GET("/", HelloController)
	e.GET("/user", user.GetUser)
	e.POST("/user", user.CreateUser)
	e.GET("/users", user.SearchUserController)
	e.GET("/users/:id", user.GetUserController)
	// run server
	e.Start(":8000")
}

// response string
func HelloController(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World")
}

// response json
func (u User) GetUser(c echo.Context) error {
	return c.JSON(http.StatusOK, u)
}

// url params
func (u *User) GetUserController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	if u.Id == id {
		return c.JSON(http.StatusOK, map[string]interface{}{"user": u})
	}
	return c.JSON(http.StatusNotFound, "id not found")
}

// query params
func (u User) SearchUserController(c echo.Context) error {
	search := c.QueryParam("search")
	if search == u.Name {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"search": search,
			"result": u,
		})
	}
	return c.JSON(http.StatusNotFound, "search not found")
}

// form
func (u *User) CreateUser(c echo.Context) error {
	u.Name = c.FormValue("name")
	u.Email = c.FormValue("email")

	// kita juga bisa memasukan formvalue sekaligus dg binding
	// c.Bind(&u)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Success create user",
		"user":    u,
	})
}
