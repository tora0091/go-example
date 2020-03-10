package service

import (
	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/entity"
)

type UsersService interface {
	GetUsersParam(c *gin.Context) (*entity.User, error)
}

type usersService struct{}

func NewUsersService() UsersService {
	return &usersService{}
}

func (*usersService) GetUsersParam(c *gin.Context) (*entity.User, error) {
	var user entity.User

	err := c.BindJSON(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
