package database

import (
	"time"
	"todo_app/cmd/domain"
)

type TodoRepository struct {
	SqlHandler SqlHandler
}

func (repository *TodoRepository) FindAllWithCategory() (todos domain.Todos, err error) {
	rows, err := repository.SqlHandler.Query(queryForFindAllWithCategory)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var description string
		var expiration string
		var categoryId int
		var todoCreatedAt time.Time
		var todoUpdatedAt time.Time
		var categoryName string
		var categoryCreatedAt time.Time
		var categoryUpdatedAt time.Time
		if err = rows.Scan(&id, &title, &description, &expiration, &categoryId, &todoCreatedAt, &todoUpdatedAt, &categoryName, &categoryCreatedAt, &categoryUpdatedAt); err != nil {
			return
		}

		category := domain.Category{
			ID:   categoryId,
			Name: categoryName,
			Base: domain.Base{
				CreatedAt: categoryCreatedAt,
				UpdatedAt: categoryUpdatedAt,
			},
		}
		todo := domain.Todo{
			ID:          id,
			Title:       title,
			Description: description,
			Expiration:  expiration,
			CategoryId:  categoryId,
			Category:    category,
			Base: domain.Base{
				CreatedAt: todoCreatedAt,
				UpdatedAt: todoUpdatedAt,
			},
		}
		todos = append(todos, todo)
	}
	return
}

func (repository *TodoRepository) FindByIdWithCategory(id int) (todo domain.Todo, err error) {
	row := repository.SqlHandler.QueryRow(queryForFindByIdWithCategory, id)
	if err != nil {
		return
	}

	var title string
	var description string
	var expiration string
	var categoryId int
	var todoCreatedAt time.Time
	var todoUpdatedAt time.Time
	var categoryName string
	var categoryCreatedAt time.Time
	var categoryUpdatedAt time.Time
	if err = row.Scan(&id, &title, &description, &expiration, &categoryId, &todoCreatedAt, &todoUpdatedAt, &categoryName, &categoryCreatedAt, &categoryUpdatedAt); err != nil {
		return
	}

	category := domain.Category{
		ID:   categoryId,
		Name: categoryName,
		Base: domain.Base{
			CreatedAt: categoryCreatedAt,
			UpdatedAt: categoryUpdatedAt,
		},
	}
	todo = domain.Todo{
		ID:          id,
		Title:       title,
		Description: description,
		Expiration:  expiration,
		CategoryId:  categoryId,
		Category:    category,
		Base: domain.Base{
			CreatedAt: todoCreatedAt,
			UpdatedAt: todoUpdatedAt,
		},
	}
	return
}

func (repository *TodoRepository) StoreTodo(todo *domain.Todo) (id int, err error) {
	result, err := repository.SqlHandler.Exec(queryForCreateTodo, todo.Title, todo.Description, todo.Expiration, todo.CategoryId, todo.CreatedAt, todo.UpdatedAt)
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
	_, err = repository.SqlHandler.Exec(queryForUpdateTodo, todo.Title, todo.Description, todo.Expiration, todo.CategoryId, todo.CreatedAt, todo.UpdatedAt, todo.ID)
	return
}

func (repository *TodoRepository) DeleteTodo(todo *domain.Todo) (err error) {
	_, err = repository.SqlHandler.Exec(queryForDeleteTodo, todo.DeletedAt, todo.ID)
	return
}
