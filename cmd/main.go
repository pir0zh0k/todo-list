package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
    var todos []string
    var command string
    reader := bufio.NewReader(os.Stdin)

    for {
        fmt.Print("Введите команду: \n 1. Создать задачу \n 2. Показать задачи \n 3. Выход \n")
        command, _ = reader.ReadString('\n')
        command = strings.TrimSpace(command)

        switch command {
        case "1":
            addTodo(&todos, reader)
        case "2":
            showTodos(todos)
        case "3":
            return
        default:
            fmt.Println("Неизвестная команда")
        }
    }
}

func addTodo(todos *[]string, reader *bufio.Reader) {
    fmt.Println("Введите текст задачи:")

    text, err := reader.ReadString('\n')
    if err != nil {
        panic(err)
    }

    text = strings.TrimSpace(text)

    *todos = append(*todos, text)
}

func showTodos(todos []string) {
    for i, todo := range todos {
        fmt.Println(i+1, todo)
    }
}