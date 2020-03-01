package handler

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"

	"go-example/restfulapi/model"
	"go-example/restfulapi/responsebody"
)

// UsersHandler ,
// curl -v -X GET http://localhost:8080/ or curl -v -X GET http://localhost:8080/users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := model.GetUserList()
	if err != nil {
		responsebody.StatusInternalServerError(w, err.Error())
		return
	}
	responsebody.StatusOKWithUsers(w, users)
}

func UserForUserIdHandler(w http.ResponseWriter, r *http.Request) {
	// vars := mux.Vars(r)
	// userId, err := strconv.Atoi(vars["user_id"])
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// if userId >= 0 && userId < len(UserList) {
	// 	err := json.NewEncoder(w).Encode(UserList[userId])
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 	}
	// }
	// log.Printf("not found user_id: %d\n", userId)
}

// CreateUserHandler ,
// curl -v -X POST -H "Content-type:application/json" -d '{"name":"andy", "job":"profession", "email":"andy@yahoo.com"}' http://localhost:8080/user
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responsebody.StatusBadRequest(w, "")
		return
	}

	if r.Header.Get("Content-type") != "application/json" {
		responsebody.StatusBadRequest(w, "")
		return
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		responsebody.StatusInternalServerError(w, err.Error())
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		responsebody.StatusInternalServerError(w, err.Error())
		return
	}

	user := model.NewUser()
	err = json.Unmarshal(body[:length], user)
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
