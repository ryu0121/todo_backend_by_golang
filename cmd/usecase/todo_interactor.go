package usecase

import (
	"strconv"
	"todo_app/cmd/domain"
)

type TodoInteractor struct {
	TodoRepository TodoRepository
}

func (interactor *TodoInteractor) Todos() (todos domain.Todos, err error) {
	todos, err = interactor.TodoRepository.FindAllWithCategory()
	return
}

func (interactor *TodoInteractor) TodoById(id string) (todo domain.Todo, err error) {
	idInt, err := strconv.Atoi(id)
	// ここでreturnした場合、todoはdomain.Todoのゼロ値(構造体の場合はフィールドが全てゼロ値の構造体)が入る
	if err != nil {
		return
	}
	todo, err = interactor.TodoRepository.FindByIdWithCategory(idInt)
	return
}

func (interactor *TodoInteractor) AddTodo(todo *domain.Todo) (err error) {
	id, err := interactor.TodoRepository.StoreTodo(todo)
	if err != nil {
		return
	}
	todo.ID = id
	return
}

func (interactor *TodoInteractor) UpdateTodo(todo *domain.Todo) (err error) {
	err = interactor.TodoRepository.UpdateTodo(todo)
	if err != nil {
		return
	}
	return
}

func (interactor *TodoInteractor) DeleteTodo(todo *domain.Todo) (err error) {
	err = interactor.TodoRepository.DeleteTodo(todo)
	if err != nil {
		return
	}
	return
}
