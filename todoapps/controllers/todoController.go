package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-example/todoapps/database"
	"go-example/todoapps/jsons"
)

func Todos(c *gin.Context) {
	db := database.GetDbConnection()
	defer db.Close()

	var todos jsons.Todos
	db.Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func Todo(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id is not numeric")
		return
	}

	db := database.GetDbConnection()
	defer db.Close()

	var todo jsons.Todo
	recordNotFound := db.First(&todo, id).RecordNotFound()
	if recordNotFound {
		c.JSON(http.StatusBadRequest, "record not found")
		return
	}
	c.JSON(http.StatusOK, todo)
}

func CreateTodo(c *gin.Context) {

}

func UpdateTodo(c *gin.Context) {

}

func DeleteTodo(c *gin.Context) {

}
