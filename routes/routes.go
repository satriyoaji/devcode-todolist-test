package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/satriyoaji/todolist-challenge-go/controllers"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()

	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, "Health Check!")
	})

	e.GET("api/activity-groups", controllers.FetchActivities)
	e.POST("api/activity-groups", controllers.StoreActivity)
	e.PUT("api/activity-groups/:id", controllers.UpdateAcitity)
	e.DELETE("api/activity-groups/:id", controllers.DeleteActivity)

	e.GET("api/generate-hash/:password", controllers.GenerateHashPassword)
	e.POST("api/login", controllers.ActionLogin)

	return e
}
