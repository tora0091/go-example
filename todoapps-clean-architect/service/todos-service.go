package service

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/entity"
)

type TodosService interface {
	GetIDParam(c *gin.Context) (int, error)
	GetRequestParam(c *gin.Context) (*entity.Todo, error)
	Validator(todo *entity.Todo) error
}

type todosService struct{}

func NewTodosService() TodosService {
	return &todosService{}
}

func (*todosService) GetIDParam(c *gin.Context) (int, error) {
	param := c.Param("id")
	id, err := strconv.Atoi(param)
	if err != nil {
		return id, err
	}
	return id, nil
}

func (*todosService) GetRequestParam(c *gin.Context) (*entity.Todo, error) {
	var todo entity.Todo

	err := c.BindJSON(&todo)
	if err != nil {
		return nil, err
	}
	return &todo, nil
}

func (*todosService) Validator(todo *entity.Todo) error {
	if todo.Title == "" {
		return fmt.Errorf("title is required")
	}
	return nil
}
