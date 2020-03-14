package router

import (
	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/controllers"
	"go-example/todoapps-clean-architect/repository"
	"go-example/todoapps-clean-architect/service"
)

var (
	todosService    service.TodosService        = service.NewTodosService()
	todosRepository repository.TodosRepository  = repository.NewTodosRepository()
	todosController controllers.TodosController = controllers.NewTodosController(todosRepository, todosService)

	usersService    service.UsersService        = service.NewUsersService()
	usersRepository repository.UsersRepository  = repository.NewUsersRepository()
	usersController controllers.UsersController = controllers.NewUsersController(usersRepository, usersService)
)

func InitRouter() *gin.Engine {
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

	return router
}
