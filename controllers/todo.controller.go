package controllers

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/todolist-challenge-go/constants"
	"github.com/satriyoaji/todolist-challenge-go/dto"
	"github.com/satriyoaji/todolist-challenge-go/models"
	"github.com/satriyoaji/todolist-challenge-go/repositories"
	"net/http"
	"strconv"
)

func FetchTodos(c echo.Context) error {
	queryParams := c.QueryParams()

	result, err := repositories.GetAllTodos(queryParams)
	if err != nil {
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	return c.JSON(http.StatusOK, result)
}

func FetchOneTodo(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}
	result, err := repositories.GetTodoByID(int_id)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	return c.JSON(http.StatusOK, result)
}

func StoreTodo(c echo.Context) error {
	var payload dto.CreateTodoPayload

	if err := c.Bind(&payload); err != nil {
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}
	if err := c.Validate(&payload); err != nil {
		return ReturnFirstErrorValidation(c, http.StatusBadRequest, constants.BadRequestStatus, err)
	}

	if payload.IsActive == nil {
		trueValue := true
		payload.IsActive = &trueValue
	}
	result, err := repositories.CreateTodo(payload.Title, payload.ActivityGroupID, *payload.IsActive)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func UpdateTodo(c echo.Context) error {
	var payload dto.UpdateTodoPayload

	if err := c.Bind(&payload); err != nil {
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}
	if err := c.Validate(&payload); err != nil {
		return ReturnFirstErrorValidation(c, http.StatusBadRequest, constants.BadRequestStatus, err)
	}

	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	resultFind, err := repositories.GetTodoByID(int_id)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, resultFind)
		}
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	todoFind := resultFind.Data.(models.Todo)
	if payload.IsActive == nil {
		payload.IsActive = &todoFind.IsActive
	}
	if payload.Title == "" {
		payload.Title = todoFind.Title
	}
	if payload.Priority == "" {
		payload.Priority = "very-high"
	}
	result, err := repositories.UpdateTodo(int_id, *payload.IsActive, payload.Title, payload.Priority)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	return c.JSON(http.StatusOK, result)
}

func DeleteTodo(c echo.Context) error {
	id := c.Param("id")
	int_id, err := strconv.Atoi(id)
	if err != nil {
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	result, err := repositories.DeleteTodoByID(int_id)
	if err != nil {
		if err.Error() == "not_found" {
			return c.JSON(http.StatusNotFound, result)
		}
		return ReturnErrorResponse(c, http.StatusInternalServerError, constants.ServerErrorStatus, err)
	}

	return c.JSON(http.StatusOK, result)
}
