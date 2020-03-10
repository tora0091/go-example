package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go-example/todoapps-clean-architect/controllers"
	"go-example/todoapps-clean-architect/database"
)

var (
	todosController controllers.TodosController = controllers.NewTodosController()
	usersController controllers.UsersController = controllers.NewUsersController()
	dbDispater      database.Database           = database.NewDatabase()
)

func init() {
	dbDispater.InitDatabase()
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", usersController.Users)
		v1.POST("/user", usersController.CreateUser)
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/todos", todosController.Todos)
		v2.GET("/todo/:id", todosController.Todo)
		v2.POST("/todo", todosController.CreateTodo)
		v2.PUT("/todo/:id", todosController.UpdateTodo)
		v2.DELETE("/todo/:id", todosController.DeleteTodo)
	}
	router.Run(":8080")
}
