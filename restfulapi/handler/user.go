package handler

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"go-example/restfulapi/model"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	// err := json.NewEncoder(w).Encode(UserList)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
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

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)
	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user model.User
	err = json.Unmarshal(body[:length], &user)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = user.CreateValidator()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = user.CreateUser()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(nil)
}
