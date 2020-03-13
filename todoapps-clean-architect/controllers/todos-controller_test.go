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

func Test_TodosControllerTodos(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}

		todos := entity.Todos{
			entity.Todo{Title: "first title", Completed: false},
			entity.Todo{Title: "second title", Completed: true},
			entity.Todo{Title: "thred title", Completed: false},
		}

		r.Mock.On("FindAll").Return(&todos, nil)
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		controller.Todos(c)
		assert.Equal(t, w.Code, http.StatusOK)

		var d jsons.JSONStatusOKWithDataResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusOK)

		dataSlice := reflect.ValueOf(d.Data)
		assert.Equal(t, dataSlice.Len(), len(todos))

		for i := 0; i < dataSlice.Len(); i++ {
			itre := dataSlice.Index(i).Elem().MapRange()
			var todo entity.Todo
			for itre.Next() {
				switch itre.Key().String() {
				case "Title":
					todo.Title = itre.Value().Interface().(string)
				case "Completed":
					todo.Completed = itre.Value().Interface().(bool)
				}
			}
			assert.Equal(t, todo.Title, todos[i].Title)
			assert.Equal(t, todo.Completed, todos[i].Completed)
		}
	})

	t.Run("ng pattern (status:500)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}

		message := "find all error"
		r.Mock.On("FindAll").Return(&entity.Todos{}, fmt.Errorf(message))
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		controller.Todos(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, message)
	})
}

func Test_TodosControllerTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}

		todo := entity.Todo{Title: "first title", Completed: false}

		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		s.Mock.On("GetIDParam", c).Return(1, nil)
		r.Mock.On("FindByID", 1).Return(&todo, nil)

		controller.Todo(c)
		assert.Equal(t, w.Code, http.StatusOK)

		var d jsons.JSONStatusOKWithDataResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusOK)

		dataSlice := reflect.ValueOf(d.Data)

		keys := dataSlice.MapKeys()
		for _, key := range keys {
			value := dataSlice.MapIndex(key)
			switch value.String() {
			case "Title":
				assert.Equal(t, value, todo.Title)
			case "Completed":
				assert.Equal(t, value, todo.Completed)
			}
		}
	})

	t.Run("ng pattern (id param is not found)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		todo := entity.Todo{Title: "first title", Completed: false}
		errorMessage := "id param is not found"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		s.Mock.On("GetIDParam", c).Return(0, fmt.Errorf(errorMessage))
		r.Mock.On("FindByID", 1).Return(&todo, nil)

		controller.Todo(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (id is not found)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		todo := entity.Todo{Title: "first title", Completed: false}
		errorMessage := "id is not found"

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("GET", "/", body)

		s.Mock.On("GetIDParam", c).Return(1, nil)
		r.Mock.On("FindByID", 1).Return(&todo, fmt.Errorf(errorMessage))

		controller.Todo(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
		assert.Equal(t, d.Message, errorMessage)
	})

}

func Test_TodosControllerCreateTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		todo := entity.Todo{}

		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("Save", &todo).Return(&todo, nil)

		controller.CreateTodo(c)
		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("ng pattern (id param is not found)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		todo := entity.Todo{}

		errorMessage := "id param is not found"
		s.Mock.On("GetRequestParam", c).Return(&todo, fmt.Errorf(errorMessage))
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("Save", &todo).Return(&todo, nil)

		controller.CreateTodo(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (validation error)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		todo := entity.Todo{}

		errorMessage := "title is required"
		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(fmt.Errorf(errorMessage))
		r.Mock.On("Save", &todo).Return(&todo, nil)

		controller.CreateTodo(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (save error)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("POST", "/", body)

		todo := entity.Todo{}

		errorMessage := "save error"
		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("Save", &todo).Return(&todo, fmt.Errorf(errorMessage))

		controller.CreateTodo(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})
}

func Test_TodosControllerUpdateTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("PUT", "/", body)

		todo := entity.Todo{}

		s.Mock.On("GetIDParam", c).Return(1, nil)
		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("UpdateByID", 1, &todo).Return(&todo, nil)

		controller.UpdateTodo(c)
		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("ng pattern (id param is not found)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("PUT", "/", body)

		todo := entity.Todo{}

		errorMessage := "id param is not found"
		s.Mock.On("GetIDParam", c).Return(1, fmt.Errorf(errorMessage))
		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("UpdateByID", 1, &todo).Return(&todo, nil)

		controller.UpdateTodo(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (request param is not found)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("PUT", "/", body)

		todo := entity.Todo{}

		errorMessage := "request param is not found"
		s.Mock.On("GetIDParam", c).Return(1, nil)
		s.Mock.On("GetRequestParam", c).Return(&todo, fmt.Errorf(errorMessage))
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("UpdateByID", 1, &todo).Return(&todo, nil)

		controller.UpdateTodo(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (validation error)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("PUT", "/", body)

		todo := entity.Todo{}

		errorMessage := "title is required"
		s.Mock.On("GetIDParam", c).Return(1, nil)
		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(fmt.Errorf(errorMessage))
		r.Mock.On("UpdateByID", 1, &todo).Return(&todo, nil)

		controller.UpdateTodo(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (update error)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("PUT", "/", body)

		todo := entity.Todo{}

		errorMessage := "update error"
		s.Mock.On("GetIDParam", c).Return(1, nil)
		s.Mock.On("GetRequestParam", c).Return(&todo, nil)
		s.Mock.On("Validator", &todo).Return(nil)
		r.Mock.On("UpdateByID", 1, &todo).Return(&todo, fmt.Errorf(errorMessage))

		controller.UpdateTodo(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})
}

func Test_TodosControllerDeleteTodo(t *testing.T) {
	gin.SetMode(gin.TestMode)

	t.Run("ok pattern", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("DELETE", "/", body)

		todo := entity.Todo{}

		s.Mock.On("GetIDParam", c).Return(1, nil)
		r.Mock.On("DeleteByID", 1).Return(&todo, nil)

		controller.DeleteTodo(c)
		assert.Equal(t, w.Code, http.StatusOK)
	})

	t.Run("ng pattern (id param is not found)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("DELETE", "/", body)

		todo := entity.Todo{}

		errorMessage := "id param is not found"
		s.Mock.On("GetIDParam", c).Return(1, fmt.Errorf(errorMessage))
		r.Mock.On("DeleteByID", 1).Return(&todo, nil)

		controller.DeleteTodo(c)
		assert.Equal(t, w.Code, http.StatusBadRequest)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusBadRequest)
		assert.Equal(t, d.Message, errorMessage)
	})

	t.Run("ng pattern (delete error)", func(t *testing.T) {
		r := controllers.MockTodosRepository{}
		s := controllers.MockTodosService{}
		controller := NewTodosController(&r, &s)

		w := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(w)

		body := bytes.NewBufferString(``)
		c.Request, _ = http.NewRequest("DELETE", "/", body)

		todo := entity.Todo{}

		errorMessage := "delete error"
		s.Mock.On("GetIDParam", c).Return(1, nil)
		r.Mock.On("DeleteByID", 1).Return(&todo, fmt.Errorf(errorMessage))

		controller.DeleteTodo(c)
		assert.Equal(t, w.Code, http.StatusInternalServerError)

		var d jsons.JSONErrorResponse
		json.NewDecoder(w.Body).Decode(&d)
		assert.Equal(t, d.Status, http.StatusInternalServerError)
		assert.Equal(t, d.Message, errorMessage)
	})
}
