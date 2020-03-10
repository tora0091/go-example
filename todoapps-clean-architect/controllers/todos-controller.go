package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/jsons"
	"go-example/todoapps-clean-architect/repository"
	"go-example/todoapps-clean-architect/service"
)

var (
	todosRepository repository.TodosRepository = repository.NewTodosRepository()
	todosService    service.TodosService       = service.NewTodosService()
)

// curl -v -X GET http://localhost:8080/api/v2/todos
func Todos(c *gin.Context) {
	todos, err := todosRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: todos})
}

// curl -v -X GET http://localhost:8080/api/v2/todo/3
func Todo(c *gin.Context) {
	id, err := todosService.GetIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	todo, err := todosRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: todo})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"title": "hello world", "completed":false}' http://localhost:8080/api/v2/todo
func CreateTodo(c *gin.Context) {
	todo, err := todosService.GetRequestParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = todosService.Validator(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	now := time.Now()
	todo.CreatedAt = now
	todo.UpdatedAt = now

	_, err = todosRepository.Save(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}

// curl -v -X PUT -H "Content-type: application/json" -d '{"title": "hello world sample", "completed":false}' http://localhost:8080/api/v2/todo/3
func UpdateTodo(c *gin.Context) {
	id, err := todosService.GetIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	updateData, err := todosService.GetRequestParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = todosService.Validator(updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	updateData.UpdatedAt = time.Now()

	todo, err := todosRepository.UpdateByID(id, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: todo})
}

// curl -v -X DELETE http://localhost:8080/api/v2/todo/8
func DeleteTodo(c *gin.Context) {
	id, err := todosService.GetIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	_, err = todosRepository.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return

	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}
