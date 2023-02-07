package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/satriyoaji/todolist-challenge-go/config"
	"github.com/satriyoaji/todolist-challenge-go/database"
	"github.com/satriyoaji/todolist-challenge-go/routes"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (validator *CustomValidator) Validate(i interface{}) error {
	return validator.validator.Struct(i)
}

func main() {
	database.Init()

	e := routes.Init()
	e.Validator = &CustomValidator{validator: validator.New()}

	e.Logger.Fatal(e.Start(":" + config.GoDotEnvVariable("PORT")))
}
