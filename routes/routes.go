package routes

import (
	"github.com/labstack/echo/v4"
	"lessonProgram/handlers"
)

func InitRoutes(e *echo.Echo) {
	e.GET("/", handlers.Welcome)

	e.GET("/students", handlers.GetStudents)
	e.GET("/student/:id", handlers.EditStudent)
	e.POST("/student/:id", handlers.UpdateStudent)
	e.POST("/student/delete/:id", handlers.DeleteStudent)
	e.GET("/student/add", handlers.CreateStudent)
	e.POST("/student/add", handlers.SaveStudent)

	e.GET("/todos", handlers.GetTodos)
	e.GET("/todo/:id", handlers.EditTodo)
	e.POST("/todo/:id", handlers.UpdateTodo)
	e.POST("/todo/delete/:id", handlers.DeleteTodo)
	e.GET("/todo/add", handlers.CreateTodo)
	e.POST("/todo/add", handlers.SaveTodo)
}
