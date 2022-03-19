package infrastructure

import (
	"todo_app/cmd/interfaces/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() {
	router := echo.New()

	router.Use(middleware.CORS())
	router.Use(middleware.Logger())
	router.Use(middleware.Recover())

	todosController := controllers.NewTodosController(NewHandler())

	router.GET("/todos", func(c echo.Context) (err error) {
		err = todosController.Index(c)
		return
	})
	router.GET("/todos/:id", func(c echo.Context) (err error) {
		err = todosController.Show(c)
		return
	})
	router.POST("/todos", func(c echo.Context) (err error) {
		err = todosController.Create(c)
		return
	})
	router.PUT("/todos/:id", func(c echo.Context) (err error) {
		err = todosController.Update(c)
		return
	})
	router.DELETE("/todos/:id", func(c echo.Context) (err error) {
		err = todosController.Destroy(c)
		return
	})

	router.Logger.Fatal(router.Start(":8080"))
}
