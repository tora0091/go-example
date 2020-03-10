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
	usersRepository repository.UsersRepository = repository.NewUsersRepository()
	usersService    service.UsersService       = service.NewUsersService()
)

// curl -v -X GET http://localhost:8080/api/v1/users
func Users(c *gin.Context) {
	users, err := usersRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: users})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"name": "Alex Murer", "email": "alex@example.com", "address": "Los Angeles, USA", "job": "artist"}' http://localhost:8080/api/v1/user
func CreateUser(c *gin.Context) {
	user, err := usersService.GetUsersParam(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	errs := usersService.Validator(user)
	if errs != nil {
		c.JSON(http.StatusBadRequest, jsons.JSONErrorResponse{Status: http.StatusBadRequest, Message: errs})
		return
	}

	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	// todo: validation

	_, err = usersRepository.Save(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}
