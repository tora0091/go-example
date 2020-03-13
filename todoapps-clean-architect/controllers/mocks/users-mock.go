package controllers

import (
	"go-example/todoapps-clean-architect/entity"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockUsersRepository struct {
	mock.Mock
}

func (mock *MockUsersRepository) Save(user *entity.User) (*entity.User, error) {
	args := mock.Called(user)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (mock *MockUsersRepository) FindAll() (*entity.Users, error) {
	args := mock.Called()
	return args.Get(0).(*entity.Users), args.Error(1)
}

type MockUsersService struct {
	mock.Mock
}

func (mock *MockUsersService) GetUsersParam(c *gin.Context) (*entity.User, error) {
	args := mock.Called(c)
	return args.Get(0).(*entity.User), args.Error(1)
}

func (mock *MockUsersService) Validator(user *entity.User) []string {
	args := mock.Called(user)
	if _, ok := args.Get(0).([]string); ok {
		return args.Get(0).([]string)
	}
	return nil
}
