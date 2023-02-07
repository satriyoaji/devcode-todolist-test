package controllers

import (
	"github.com/labstack/echo/v4"
)

func ReturnErrorResponse(c echo.Context, code int, status string, err error) error {
	return c.JSON(code, map[string]interface{}{
		"status":  status,
		"message": err.Error(),
		"data":    map[string]interface{}{},
	})
}
