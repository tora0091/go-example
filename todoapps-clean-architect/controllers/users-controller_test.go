package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	controllers "go-example/todoapps-clean-architect/controllers/mocks"
	"go-example/todoapps-clean-architect/entity"
	"go-example/todoapps-clean-architect/jsons"
)

func Test_UsersControllerUsers(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockUsersRepository{}
		s := controllers.MockUsersService{}

		users := entity.Users{
			entity.User{Name: "John", Email: "john@exmaple.com", Address: "LA", Job: "artist"},
			entity.User{Name: "Fumiji", Email: "fumifumi@exmaple.com", Address: "NY", Job: "fumifumi"},
		}

		r.Mock.On("FindAll").Return(&users, nil)
		controller := NewUsersController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		controller.Users(c)
		assert.Equal(t, w.Code, http.StatusOK)

		var d jsons.JSONStatusOKWithDataResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusOK)

		dataSlice := reflect.ValueOf(d.Data)
		assert.Equal(t, dataSlice.Len(), len(users))

		for i := 0; i < dataSlice.Len(); i++ {
			itre := dataSlice.Index(i).Elem().MapRange()
			var user entity.User
			for itre.Next() {
				switch itre.Key().String() {
				case "Name":
					user.Name = itre.Value().Interface().(string)
				case "Email":
					user.Email = itre.Value().Interface().(string)
				case "Address":
					user.Address = itre.Value().Interface().(string)
				case "Job":
					user.Job = itre.Value().Interface().(string)
				}
			}
			assert.Equal(t, user.Name, users[i].Name)
			assert.Equal(t, user.Email, users[i].Email)
			assert.Equal(t, user.Address, users[i].Address)
			assert.Equal(t, user.Job, users[i].Job)
		}
	})

	t.Run("ng pattern (find all error)", func(t *testing.T) {
		r := controllers.MockUsersRepository{}
		s := controllers.MockUsersService{}

		users := entity.Users{
			entity.User{Name: "John", Email: "john@exmaple.com", Address: "LA", Job: "artist"},
			entity.User{Name: "Fumiji", Email: "fumifumi@exmaple.com", Address: "NY", Job: "fumifumi"},
		}

		errorMessage := "find all error"
		r.Mock.On("FindAll").Return(&users, fmt.Errorf(errorMessage))
		controller := NewUsersController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		controller.Users(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})
}

func Test_UsersControllerCreateUser(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockUsersRepository{}
		s := controllers.MockUsersService{}

		controller := NewUsersController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		user := entity.User{}
		s.Mock.On("GetUsersParam", c).Return(&user, nil)
		s.Mock.On("Validator", &user).Return(nil)
		r.Mock.On("Save", &user).Return(&user, nil)

		controller.CreateUser(c)
		assert.Equal(t, w.Code, http.StatusOK)

		var d jsons.JSONStatusOKWithDataResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusOK)
	})

	t.Run("ng pattern (user param is not found)", func(t *testing.T) {
		r := controllers.MockUsersRepository{}
		s := controllers.MockUsersService{}

		controller := NewUsersController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		errorMessage := "user param is not found"
		user := entity.User{}
		s.Mock.On("GetUsersParam", c).Return(&user, fmt.Errorf(errorMessage))
		s.Mock.On("Validator", &user).Return(nil)
		r.Mock.On("Save", &user).Return(&user, nil)

		controller.CreateUser(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (validation error)", func(t *testing.T) {
		r := controllers.MockUsersRepository{}
		s := controllers.MockUsersService{}

		controller := NewUsersController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		errorMessages := []string{
			"name is required",
			"email is required",
			"address is required",
			"job is required",
		}
		user := entity.User{}
		s.Mock.On("GetUsersParam", c).Return(&user, nil)
		s.Mock.On("Validator", &user).Return(errorMessages)
		r.Mock.On("Save", &user).Return(&user, nil)

		controller.CreateUser(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
	})

	t.Run("ng pattern (save error)", func(t *testing.T) {
		r := controllers.MockUsersRepository{}
		s := controllers.MockUsersService{}

		controller := NewUsersController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		errorMessage := "save error"
		user := entity.User{}
		s.Mock.On("GetUsersParam", c).Return(&user, nil)
		s.Mock.On("Validator", &user).Return(nil)
		r.Mock.On("Save", &user).Return(&user, fmt.Errorf(errorMessage))

		controller.CreateUser(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})
}
