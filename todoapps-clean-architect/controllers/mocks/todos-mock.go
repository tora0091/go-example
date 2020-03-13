package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"

	"go-example/todoapps-clean-architect/entity"
)

type MockTodosRepository struct {
	mock.Mock
}

func (mock *MockTodosRepository) Save(todo *entity.Todo) (*entity.Todo, error) {
	args := mock.Called(todo)
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (mock *MockTodosRepository) FindAll() (*entity.Todos, error) {
	args := mock.Called()
	return args.Get(0).(*entity.Todos), args.Error(1)
}

func (mock *MockTodosRepository) FindByID(id int) (*entity.Todo, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (mock *MockTodosRepository) UpdateByID(id int, updateData *entity.Todo) (*entity.Todo, error) {
	args := mock.Called(id, updateData)
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (mock *MockTodosRepository) DeleteByID(id int) (*entity.Todo, error) {
	args := mock.Called(id)
	return args.Get(0).(*entity.Todo), args.Error(1)
}

type MockTodosService struct {
	mock.Mock
}

func (mock *MockTodosService) GetIDParam(c *gin.Context) (int, error) {
	args := mock.Called(c)
	return args.Get(0).(int), args.Error(1)
}

func (mock *MockTodosService) GetRequestParam(c *gin.Context) (*entity.Todo, error) {
	args := mock.Called(c)
	return args.Get(0).(*entity.Todo), args.Error(1)
}

func (mock *MockTodosService) Validator(todo *entity.Todo) error {
	args := mock.Called(todo)
	return args.Error(0)
}
