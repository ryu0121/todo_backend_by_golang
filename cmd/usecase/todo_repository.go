package usecase

import "todo_app/cmd/domain"

type TodoRepository interface {
	FindAll() (domain.Todos, error)
	FindById(int) (domain.Todo, error)
	StoreTodo(*domain.Todo) (int, error)
	UpdateTodo(*domain.Todo) error
	DeleteTodo(*domain.Todo) error
}
