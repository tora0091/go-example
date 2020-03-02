package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"go-example/restfulapi/model"
	"go-example/restfulapi/responsebody"
)

// UsersHandler ,
// curl -v -X GET -H 'Authorization: Bearer {auth key}' http://localhost:8080/ or curl -v -X GET -H 'Authorization: Bearer {auth key}' http://localhost:8080/api/users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUserList()
	if err != nil {
		responsebody.StatusInternalServerError(w, err.Error())
		return
	}
	responsebody.StatusOKWithUsers(w, users)
}

// UserForUserIdHandler ,
// curl -v -X GET -H 'Authorization: Bearer {auth key}' http://localhost:8080/api/user/1
func UserForUserIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := model.NewUser()
	user.Id = vars["user_id"]

	target, err := user.GetUser()
	if err != nil {
		responsebody.StatusBadRequest(w, err.Error())
		return
	}
	responsebody.StatusOKWithUser(w, target)
}

// CreateUserHandler ,
// curl -v -X POST -H 'Authorization: Bearer {auth key}' -H "Content-type:application/json" -d '{"name":"andy", "job":"profession", "email":"andy@yahoo.com"}' http://localhost:8080/api/user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getJsonBody(r)
	if err != nil {
		responsebody.StatusInternalServerError(w, err.Error())
		return
	}

	err = user.CreateValidator()
	if err != nil {
		responsebody.StatusBadRequest(w, err.Error())
		return
	}

	err = user.CreateUser()
	if err != nil {
		responsebody.StatusBadRequest(w, err.Error())
		return
	}

	responsebody.StatusOK(w)
}

// DeleteUserHandler ,
// curl -v -X DELETE -H 'Authorization: Bearer {auth key}' http://localhost:8080/api/user/1
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	user := model.NewUser()
	user.Id = vars["user_id"]

	err := user.DeleteUser()
	if err != nil {
		responsebody.StatusBadRequest(w, err.Error())
		return
	}
	responsebody.StatusOK(w)
}

// UpdateUserHandler ,
// curl -v -X PUT -H 'Authorization: Bearer {auth key}' -H "Content-type:application/json" -d '{"name":"watanabe akira", "job":"actor", "email":"w.akkun92@yahoo.com"}' http://localhost:8080/api/user/10
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	user, err := getJsonBody(r)
	if err != nil {
		responsebody.StatusInternalServerError(w, err.Error())
		return
	}

	vars := mux.Vars(r)
	user.Id = vars["user_id"]

	err = user.UpdateValidator()
	if err != nil {
		responsebody.StatusBadRequest(w, err.Error())
		return
	}

	err = user.UpdateUser()
	if err != nil {
		responsebody.StatusBadRequest(w, err.Error())
		return
	}

	responsebody.StatusOKWithUser(w, *user)
}

func getJsonBody(r *http.Request) (*model.User, error) {
	if r.Header.Get("Content-type") != "application/json" {
		return nil, fmt.Errorf("missing content type")
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return nil, err
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		return nil, err
	}

	user := model.NewUser()
	err = json.Unmarshal(body[:length], user)
	if err != nil {
		return nil, err
	}
	return user, nil
}
