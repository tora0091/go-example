package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-example/todoapps-clean-architect/entity"
	"go-example/todoapps-clean-architect/jsons"
	"go-example/todoapps-clean-architect/repository"
)

var (
	usersRepository repository.UsersRepository = repository.NewUsersRepository()
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
	var user entity.User
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// todo: validation

	_, err = usersRepository.Save(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
}
