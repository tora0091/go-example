package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/jsons"
	"go-example/todoapps-clean-architect/repository"
	"go-example/todoapps-clean-architect/service"
)

type TodosController interface {
	Todos(c *gin.Context)
	Todo(c *gin.Context)
	CreateTodo(c *gin.Context)
	UpdateTodo(c *gin.Context)
	DeleteTodo(c *gin.Context)
}

type todosController struct {
	todosRepository repository.TodosRepository
	todosService    service.TodosService
}

func NewTodosController(r repository.TodosRepository, s service.TodosService) TodosController {
	return &todosController{
		todosRepository: r,
		todosService:    s,
	}
}

// curl -v -X GET http://localhost:8080/api/v2/todos
func (t *todosController) Todos(c *gin.Context) {
	todos, err := t.todosRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: todos})
}

// curl -v -X GET http://localhost:8080/api/v2/todo/3
func (t *todosController) Todo(c *gin.Context) {
	id, err := t.todosService.GetIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	todo, err := t.todosRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: todo})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"title": "hello world", "completed":false}' http://localhost:8080/api/v2/todo
func (t *todosController) CreateTodo(c *gin.Context) {
	todo, err := t.todosService.GetRequestParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = t.todosService.Validator(todo)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	now := time.Now()
	todo.CreatedAt = now
	todo.UpdatedAt = now

	_, err = t.todosRepository.Save(todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}

// curl -v -X PUT -H "Content-type: application/json" -d '{"title": "hello world sample", "completed":false}' http://localhost:8080/api/v2/todo/3
func (t *todosController) UpdateTodo(c *gin.Context) {
	id, err := t.todosService.GetIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	updateData, err := t.todosService.GetRequestParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	err = t.todosService.Validator(updateData)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
	}

	updateData.UpdatedAt = time.Now()

	todo, err := t.todosRepository.UpdateByID(id, updateData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: todo})
}

// curl -v -X DELETE http://localhost:8080/api/v2/todo/8
func (t *todosController) DeleteTodo(c *gin.Context) {
	id, err := t.todosService.GetIDParam(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
		return
	}

	_, err = t.todosRepository.DeleteByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return

	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}
