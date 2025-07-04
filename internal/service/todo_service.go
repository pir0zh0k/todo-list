package service

import (
	"errors"
	"strconv"
	"todo-list/internal/model"
)

type TodoService struct {
	todos []model.Todo
}


func NewTodoService() *TodoService {
	return &TodoService{
		todos: make([]model.Todo, 0),
	}
}

func (s *TodoService) Add(text string) {
	todo := model.Todo{
		ID: len(s.todos) + 1,
		Text: text,
		Completed: false,
	}

	s.todos = append(s.todos, todo)
}

func (s *TodoService) List() []model.Todo {
	return  s.todos
}

func (s *TodoService) Complete(idStr string) error {
	id, err := strconv.Atoi(idStr)
	
	if err != nil {
		return errors.New("не корректный ID")
	}

	for i, todo := range s.todos {
		if todo.ID == id {
			s.todos[i].Completed = true
			return nil
		}
	}

	return errors.New("задача не найдена")
}