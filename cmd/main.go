package main

import (
	"todo-list/internal/cli"
	"todo-list/internal/service"
)

func main() {
	todoService := service.NewTodoService()
	app := cli.NewCli(todoService)

	app.Run()
}
