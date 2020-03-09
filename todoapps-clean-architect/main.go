package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go-example/todoapps-clean-architect/controllers"
	"go-example/todoapps-clean-architect/database"
)

func init() {
	database.InitDatabase()
}

func main() {
	router := gin.Default()

	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", controllers.Users)
		v1.POST("/user", controllers.CreateUser)
	}

	// v2 := router.Group("/api/v2")
	// {
	// 	v2.GET("/todos", controllers.Todos)
	// 	v2.GET("/todo/:id", controllers.Todo)
	// 	v2.POST("/todo", controllers.CreateTodo)
	// 	v2.PUT("/todo/:id", controllers.UpdateTodo)
	// 	v2.DELETE("/todo/:id", controllers.DeleteTodo)
	// }
	router.Run(":8080")
}
