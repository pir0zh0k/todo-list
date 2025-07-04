package cli

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"todo-list/internal/service"
)

type CLI struct {
	service *service.TodoService
	reader *bufio.Reader
}

func NewCli(s *service.TodoService) *CLI {
	return &CLI{
		service: s,
		reader: bufio.NewReader(os.Stdin),
	}
}

func (cli *CLI) Run() {
	for {
		fmt.Print("Введите команду: \n 1. Создать задачу \n 2. Показать задачи \n 3. Выполнить задачу \n 4. Выход \n")

		command, _ := cli.reader.ReadString('\n')
		command = strings.TrimSpace(command)

		switch command {
			case "1":
				cli.addTodo()
			case "2":
				cli.showTodos()
			case "3":
				cli.completeTask()
			case "4":
				return
			default:
				fmt.Println("Неизвестная команда")
		}
	}
}

func (cli *CLI) addTodo() {
	fmt.Println("Введите текст задачи:")
    text, _ := cli.reader.ReadString('\n')
    text = strings.TrimSpace(text)

    cli.service.Add(text)
    fmt.Println("Задача добавлена.")
}

func (cli *CLI) showTodos() {
	todos := cli.service.List()

	if len(todos) == 0 {
		fmt.Println("Список пуст")
		return
	}

	for _, todo := range todos {
		status := "❌"

		if todo.Completed {
			status = "✅"
		}

		fmt.Printf("%d. %s [%s]\n", todo.ID, todo.Text, status)
	}
}

func (cli *CLI) completeTask() {
	fmt.Println("Введите ID задачи:")

	id, _ := cli.reader.ReadString('\n')
	id = strings.TrimSpace(id)

	err := cli.service.Complete(id)

	if err != nil {
		fmt.Println("Ошибка", err)
	} else {
		fmt.Println("Задача выполнена")
	}
}