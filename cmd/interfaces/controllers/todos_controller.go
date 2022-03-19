package controllers

import (
	"net/http"
	"strconv"
	"todo_app/cmd/domain"
	"todo_app/cmd/interfaces/database"
	"todo_app/cmd/usecase"
)

type TodosController struct {
	Interactor usecase.TodoInteractor
}

type Context interface {
	JSON(int, interface{}) error
	Param(name string) string
	Bind(i interface{}) error
}

func NewTodosController(sqlHandler database.SqlHandler) *TodosController {
	return &TodosController{
		Interactor: usecase.TodoInteractor{
			TodoRepository: &database.TodoRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller TodosController) Index(c Context) (err error) {
	todos, err := controller.Interactor.Todos()
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusAccepted, todos)
	return
}

func (controller TodosController) Show(c Context) (err error) {
	id := c.Param("id")
	todo, err := controller.Interactor.TodoById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusAccepted, todo)
	return
}

func (controller TodosController) Create(c Context) (err error) {
	todo := new(domain.Todo)
	if err = c.Bind(todo); err != nil {
		return
	}
	err = controller.Interactor.AddTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusAccepted, *todo)
	return
}

func (controller TodosController) Update(c Context) (err error) {
	todo := new(domain.Todo)
	if err = c.Bind(todo); err != nil {
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	todo.ID = id
	err = controller.Interactor.UpdateTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusAccepted, *todo)
	return
}

func (controller TodosController) Destroy(c Context) (err error) {
	todo := new(domain.Todo)
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	todo.ID = id

	err = controller.Interactor.DeleteTodo(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	c.JSON(http.StatusAccepted, *todo)
	return
}
