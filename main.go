package main

import (
	"github.com/satriyoaji/todolist-challenge-go/config"
	"github.com/satriyoaji/todolist-challenge-go/database"
	"github.com/satriyoaji/todolist-challenge-go/routes"
)

func main() {
	database.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":" + config.GoDotEnvVariable("PORT")))
}
