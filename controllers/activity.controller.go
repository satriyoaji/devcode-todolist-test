package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/todolist-challenge-go/repositories"
	"net/http"
	"strconv"
)

func FetchActivities(c echo.Context) error {
	result, err := repositories.GetAllActivities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchOneActivity(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	result, err := repositories.GetActivityByID(int_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreActivity(c echo.Context) error {
	title := c.FormValue("title")
	email := c.FormValue("email")

	result, err := repositories.CreateActivity(title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateActivity(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	title := c.FormValue("title")
	email := c.FormValue("email")

	result, err := repositories.UpdateActivity(int_id, title, email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteActivity(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	result, err := repositories.DeleteActivityByID(int_id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, result)
}
