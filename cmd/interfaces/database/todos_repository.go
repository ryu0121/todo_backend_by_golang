package database

import (
	"todo_app/cmd/domain"
)

type TodoRepository struct {
	SqlHandler SqlHandler
}

func (repository *TodoRepository) FindAll() (todos domain.Todos, err error) {
	rows, err := repository.SqlHandler.Query(queryForFindAllTodos)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var content string
		var checked bool
		var removed bool
		if err = rows.Scan(&id, &content, &checked, &removed); err != nil {
			return
		}

		todo := domain.Todo{
			ID:      id,
			Content: content,
			Checked: checked,
			Removed: removed,
		}
		todos = append(todos, todo)
	}
	return
}

func (repository *TodoRepository) FindById(id int) (todo domain.Todo, err error) {
	row := repository.SqlHandler.QueryRow(queryForFindById, id)
	if err != nil {
		return
	}

	var content string
	var checked bool
	var removed bool
	if err = row.Scan(&id, &content, &checked, &removed); err != nil {
		return
	}

	todo = domain.Todo{
		ID:      id,
		Content: content,
		Checked: checked,
		Removed: removed,
	}
	return
}

func (repository *TodoRepository) StoreTodo(todo *domain.Todo) (id int, err error) {
	result, err := repository.SqlHandler.Exec(queryForCreateTodo, todo.Content)
	if err != nil {
		return
	}

	id64, err := result.LastInsertId()
	if err != nil {
		return
	}

	id = int(id64)
	return
}

func (repository *TodoRepository) UpdateTodo(todo *domain.Todo) (err error) {
	_, err = repository.SqlHandler.Exec(queryForUpdateTodo, todo.Content, todo.Checked, todo.Removed, todo.ID)
	return
}

func (repository *TodoRepository) DeleteTodo(todo *domain.Todo) (err error) {
	_, err = repository.SqlHandler.Exec(queryForDeleteTodo, todo.ID)
	return
}
