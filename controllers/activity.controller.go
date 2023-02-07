package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/todolist-challenge-go/constants"
	"github.com/satriyoaji/todolist-challenge-go/dto"
	"github.com/satriyoaji/todolist-challenge-go/repositories"
	"net/http"
	"strconv"
)

func FetchActivities(c echo.Context) error {
	result, err := repositories.GetAllActivities()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	return c.JSON(http.StatusOK, result)
}

func FetchOneActivity(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}
	result, err := repositories.GetActivityByID(int_id)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	return c.JSON(http.StatusOK, result)
}

func StoreActivity(c echo.Context) error {
	var payload dto.CreateActivityPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}
	if err := c.Validate(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  constants.BadRequestStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	result, err := repositories.CreateActivity(payload.Title, payload.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	return c.JSON(http.StatusOK, result)
}

func UpdateActivity(c echo.Context) error {
	var payload dto.UpdateActivityPayload

	if err := c.Bind(&payload); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}
	if err := c.Validate(&payload); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"status":  constants.BadRequestStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	result, err := repositories.UpdateActivity(int_id, payload.Title)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteActivity(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	result, err := repositories.DeleteActivityByID(int_id)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"status":  constants.ServerErrorStatus,
			"message": err.Error(),
			"data":    "{}",
		})
	}

	return c.JSON(http.StatusOK, result)
}
