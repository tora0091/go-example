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

func Test_TodosServiceValidator(t *testing.T) {
	t.Run("ok pattern", func(t *testing.T) {
		todo := entity.Todo{}
		service := NewTodosService()

		todo.Title = "Today is good day"
		actual := service.Validator(&todo)

		assert.Equal(t, nil, actual)
	})

	t.Run("ng pattern (title is blank)", func(t *testing.T) {
		todo := entity.Todo{}
		service := NewTodosService()

		todo.Title = ""
		actual := service.Validator(&todo)
		expected := "title is required"

		assert.Equal(t, expected, actual.Error())
	})
}

func Test_TodosServiceGetIDParam(t *testing.T) {
	service := NewTodosService()

	t.Run("ok pattern (value is number)", func(t *testing.T) {
		var params gin.Params
		params = append(params, gin.Param{Key: "id", Value: "12345"})
		c := gin.Context{Params: params}

		num, err := service.GetIDParam(&c)

		assert.Equal(t, nil, err)
		assert.Equal(t, 12345, num)
	})

	t.Run("ng pattern (value is not number(abc))", func(t *testing.T) {
		var params gin.Params
		params = append(params, gin.Param{Key: "id", Value: "abc"})
		c := gin.Context{Params: params}

		_, err := service.GetIDParam(&c)

		assert.Equal(t, "strconv.Atoi: parsing \"abc\": invalid syntax", err.Error())
	})

	t.Run("ng pattern (value is empty)", func(t *testing.T) {
		var params gin.Params
		params = append(params, gin.Param{Key: "id", Value: ""})
		c := gin.Context{Params: params}

		_, err := service.GetIDParam(&c)

		assert.Equal(t, "strconv.Atoi: parsing \"\": invalid syntax", err.Error())
	})
}

func Test_TodosServiceGetRequestParam(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())

		body := bytes.NewBufferString(`{"title":"This is a title", "completed":true}`)
		c.Request, _ = http.NewRequest("POST", "/", body)
		c.Request.Header.Add("Content-Type", "application/json")

		service := NewTodosService()
		entity, err := service.GetRequestParam(c)

		assert.Equal(t, err, nil)
		assert.Equal(t, entity.Title, "This is a title")
		assert.Equal(t, entity.Completed, true)
	})

	t.Run("ok pattern (no json data)", func(t *testing.T) {
		c, _ := gin.CreateTestContext(httptest.NewRecorder())

		body := bytes.NewBufferString(`{}`)
		c.Request, _ = http.NewRequest("POST", "/", body)
		c.Request.Header.Add("Content-Type", "application/json")

		service := NewTodosService()
		entity, err := service.GetRequestParam(c)

		assert.Equal(t, err, nil)
		assert.Equal(t, entity.Title, "")
		assert.Equal(t, entity.Completed, false)
	})
}
