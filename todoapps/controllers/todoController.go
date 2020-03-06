package controllers

import (
	"net/http"
	"strconv"
	"time"

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

// curl -v -X POST -H "Content-type: application/json" -d '{"title": "hello world", "completed":false}' http://localhost:8080/api/v2/todo
func CreateTodo(c *gin.Context) {
	todo := jsons.Todo{}
	now := time.Now()
	todo.CreatedAt = now
	todo.UpdatedAt = now

	err := c.BindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	db := database.GetDbConnection()
	db.NewRecord(todo)
	db.Create(&todo)
	if db.NewRecord(todo) == false {
		c.JSON(http.StatusOK, nil)
		return
	}
	c.JSON(http.StatusInternalServerError, "record not found")
}

// curl -v -X PUT -H "Content-type: application/json" -d '{"title": "hello world sample", "completed":false}' http://localhost:8080/api/v2/todo/3
func UpdateTodo(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id is not numeric")
		return
	}

	updateData := jsons.Todo{}
	updateData.UpdatedAt = time.Now()

	err = c.BindJSON(&updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	todo := jsons.Todo{}
	db := database.GetDbConnection()
	if err = db.First(&todo, id).Update(&updateData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)
}

// curl -v -X DELETE http://localhost:8080/api/v2/todo/8
func DeleteTodo(c *gin.Context) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		c.JSON(http.StatusBadRequest, "id is not numeric")
		return
	}

	todo := jsons.Todo{}
	db := database.GetDbConnection()
	if err = db.Delete(&todo, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, todo)
}
