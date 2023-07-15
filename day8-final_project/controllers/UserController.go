package controllers

import (
	"day_eight/lib/database"
	"day_eight/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func UserIndex(c echo.Context) error {
	users, e := database.GetUsers()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   users,
	})
}

func UserShow(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, e := database.FindUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"user":   user,
	})
}

func UserCreate(c echo.Context) error {
	var user models.Users
	c.Bind(&user)
	res, e := database.CreateUser(user)

	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    res,
	})
}

func UserUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var user models.Users
	c.Bind(&user)
	res, e := database.UpdateUser(id, user)
	if e != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update user",
		"user":    res,
	})
}

func UserDelete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	user, e := database.DeleteUser(id)
	if e != nil {
		return echo.NewHTTPError(http.StatusNotFound, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success delete user",
		"user":   user,
	})
}
