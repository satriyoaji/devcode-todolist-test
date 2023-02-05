package middlewares

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/satriyoaji/todolist-challenge-go/controllers"
)

var IsAuthenticated = middleware.JWTWithConfig(middleware.JWTConfig{
	SigningKey: []byte(controllers.SecretKey),
})
