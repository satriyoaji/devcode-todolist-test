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

	e.GET("activity-groups", controllers.FetchActivities)
	e.GET("activity-groups/:id", controllers.FetchOneActivity)
	e.POST("activity-groups", controllers.StoreActivity)
	e.PATCH("activity-groups/:id", controllers.UpdateActivity)
	e.DELETE("activity-groups/:id", controllers.DeleteActivity)

	e.GET("todo-items", controllers.FetchTodos)
	e.GET("todo-items/:id", controllers.FetchOneTodo)
	e.POST("todo-items", controllers.StoreTodo)
	e.PATCH("todo-items/:id", controllers.UpdateTodo)
	e.DELETE("todo-items/:id", controllers.DeleteTodo)

	//e.GET("api/generate-hash/:password", controllers.GenerateHashPassword)
	//e.POST("api/login", controllers.ActionLogin)

	return e
}
