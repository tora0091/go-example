package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"go-example/todoapps/database"
	"go-example/todoapps/jsons"
)

// curl -v -X GET http://localhost:8080/api/v1/users
func Users(c *gin.Context) {
	db := database.GetDbConnection()
	defer db.Close()

	var users jsons.Users
	db.Find(&users)
	c.JSON(http.StatusOK, jsons.JSONStatusOKWithDataResponse{Status: http.StatusOK, Data: users})
}

// curl -v -X POST -H "Content-type: application/json" -d '{"name": "Alex Murer", "email": "alex@example.com", "address": "Los Angeles, USA", "job": "artist"}' http://localhost:8080/api/v1/user
func CreateUser(c *gin.Context) {
	var user jsons.User
	now := time.Now()
	user.CreatedAt = now
	user.UpdatedAt = now

	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: err.Error()})
		return
	}

	// todo: validation

	db := database.GetDbConnection()
	db.NewRecord(user)
	db.Create(&user)
	if db.NewRecord(user) == false {
		c.JSON(http.StatusOK, jsons.JSONStatusOKResponse{Status: http.StatusOK})
		return
	}
	c.JSON(http.StatusInternalServerError, jsons.JSONErrorResponse{Status: http.StatusInternalServerError, Message: "create error"})
}
