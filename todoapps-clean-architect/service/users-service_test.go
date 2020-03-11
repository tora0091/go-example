package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"go-example/todoapps-clean-architect/entity"
)

func Test_UsersServiceValidator(t *testing.T) {
	service := NewUsersService()

	t.Run("ok pattern", func(t *testing.T) {
		users := entity.User{
			Name:    "Nick",
			Email:   "nick@example.com",
			Address: "Big Apple",
			Job:     "scientist",
		}
		result := service.Validator(&users)

		assert.Equal(t, len(result), 0)
	})

	t.Run("ng pattern", func(t *testing.T) {
		users := entity.User{}
		result := service.Validator(&users)

		assert.Equal(t, len(result), 4)
		assert.Equal(t, result[0], "name is required")
		assert.Equal(t, result[1], "email is required")
		assert.Equal(t, result[2], "address is required")
		assert.Equal(t, result[3], "job is required")
	})
}

func Test_UsersServiceGetUsersParam(t *testing.T) {
	gin.SetMode(gin.TestMode)
	service := NewUsersService()

	t.Run("ok pattern", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())

		body := bytes.NewBufferString(`{"name":"Alex", "email":"alex@example.com", "address":"New York", "job":"actor"}`)
		c.Request, _ = http.NewRequest("POST", "/", body)
		c.Request.Header.Add("Content-Type", "application/json")

		entity, err := service.GetUsersParam(c)

		assert.Equal(t, err, nil)
		assert.Equal(t, entity.Name, "Alex")
		assert.Equal(t, entity.Email, "alex@example.com")
		assert.Equal(t, entity.Address, "New York")
		assert.Equal(t, entity.Job, "actor")
	})

	t.Run("ok pattern (no json data)", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())

		body := bytes.NewBufferString(`{}`)
		c.Request, _ = http.NewRequest("POST", "/", body)
		c.Request.Header.Add("Content-Type", "application/json")

		entity, err := service.GetUsersParam(c)

		assert.Equal(t, err, nil)
		assert.Equal(t, entity.Name, "")
		assert.Equal(t, entity.Email, "")
		assert.Equal(t, entity.Address, "")
		assert.Equal(t, entity.Job, "")
	})
}
