package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/jsons"
	"go-example/todoapps-clean-architect/repository"
	"go-example/todoapps-clean-architect/service"
)

type UsersController interface {
	Users(c *gin.Context)
	CreateUser(c *gin.Context)
}

type usersController struct {
	usersRepository repository.UsersRepository
	usersService    service.UsersService
}

func NewUsersController(r repository.UsersRepository, s service.UsersService) UsersController {
	return &usersController{
		usersRepository: r,
		usersService:    s,
	}
}

// curl -v -X GET http://localhost:8080/api/v1/users
func (u *usersController) Users(c *gin.Context) {
	users, err := u.usersRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: users})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"name": "Alex Murer", "email": "alex@example.com", "address": "Los Angeles, USA", "job": "artist"}' http://localhost:8080/api/v1/user
func (u *usersController) CreateUser(c *gin.Context) {
	user, err := u.usersService.GetUsersParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	errs := u.usersService.Validator(user)
	if errs != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: errs})
		return
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	_, err = u.usersRepository.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}
