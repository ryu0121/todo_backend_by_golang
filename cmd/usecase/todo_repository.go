package usecase

import "todo_app/cmd/domain"

type TodoRepository interface {
	FindAllWithCategory() (domain.Todos, error)
	FindByIdWithCategory(int) (domain.Todo, error)
	StoreTodo(*domain.Todo) (int, error)
	UpdateTodo(*domain.Todo) error
	DeleteTodo(*domain.Todo) error
}
