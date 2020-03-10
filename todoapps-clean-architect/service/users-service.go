package service

import (
	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/entity"
)

type UsersService interface {
	GetUsersParam(c *gin.Context) (*entity.User, error)
	Validator(user *entity.User) []string
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

func (*usersService) Validator(user *entity.User) []string {
	var errs []string
	if user.Name == "" {
		errs = append(errs, "name is required")
	}
	if user.Email == "" {
		errs = append(errs, "email is required")
	}
	if user.Address == "" {
		errs = append(errs, "address is required")
	}
	if user.Job == "" {
		errs = append(errs, "job is required")
	}

	if len(errs) > 0 {
		return errs
	}
	return nil
}
