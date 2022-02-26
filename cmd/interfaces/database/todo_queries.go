package database

const (
	queryForFindAllWithCategory = `
		SELECT
			todo.id,
			todo.title,
			todo.description,
			todo.expiration,
			category.id,
			todo.created_at,
			todo.updated_at,
			category.name,
			category.created_at,
			category.updated_at
		FROM
			todos todo
		INNER JOIN categories category
			ON category.id = todo.category_id
				AND category.deleted_at IS NULL
		WHERE
			todo.deleted_at IS NULL
	`
	queryForFindByIdWithCategory = `
		SELECT
			todo.id,
			todo.title,
			todo.description,
			todo.expiration,
			category.id,
			todo.created_at,
			todo.updated_at,
			category.name,
			category.created_at,
			category.updated_at
		FROM
			todos todo
		INNER JOIN categories category
			ON category.id = todo.category_id
				AND category.deleted_at IS NULL
		WHERE
			todo.deleted_at IS NULL
			AND todo.id = ?
	`
	queryForCreateTodo = `
		INSERT INTO todos (title, description, expiration, category_id, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	queryForUpdateTodo = `
		UPDATE todos
		SET title = ?, description = ?, expiration = ?, category_id = ?, created_at = ?, updated_at = ?
		WHERE id = ?
	`

	queryForDeleteTodo = `
		UPDATE todos
		SET deleted_at = ?
		WHERE id = ?
	`
)
