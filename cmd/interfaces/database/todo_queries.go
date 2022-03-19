package database

const (
	queryForFindAllTodos = `
		SELECT
			todo.id,
			todo.content,
			todo.checked,
			todo.removed
		FROM
			todos todo
	`
	queryForFindById = `
		SELECT
			todo.id,
			todo.content,
			todo.checked,
			todo.removed
		FROM
			todos todo
		WHERE
			todo.id = ?
	`
	queryForCreateTodo = `
		INSERT INTO todos (content, checked, removed)
		VALUES (?, FALSE, FALSE)
	`

	queryForUpdateTodo = `
		UPDATE todos
		SET content = ?, checked = ?, removed = ?
		WHERE id = ?
	`

	queryForDeleteTodo = `
		DELETE FROM todos
		WHERE id = ?
	`
)
