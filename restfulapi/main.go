package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Job   string `json:"job"`
}

type Users []User

var UserList Users

func init() {
	log.SetPrefix("[server] ")

	UserList = Users{
		User{Name: "Taro Yamada", Email: "yamada@google.com", Job: "Food professer"},
		User{Name: "Kyoko Shindo", Email: "k.shindo@google.com", Job: "Athlete"},
		User{Name: "Tsukasa Kudo", Email: "kudo.t@google.com", Job: "web developer"},
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", UsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{user_id:[0-9]+}", UserForUserIdHandler).Methods(http.MethodGet)

	log.Println("start server localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	UsersHandler(w, r)
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	err := json.NewEncoder(w).Encode(UserList)
	if err != nil {
		log.Fatalln(err)
	}
}

func UserForUserIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["user_id"])
	if err != nil {
		log.Fatalln(err)
	}

	if userId >= 0 && userId < len(UserList) {
		err := json.NewEncoder(w).Encode(UserList[userId])
		if err != nil {
			log.Fatalln(err)
		}
	}
	log.Printf("not found user_id: %d\n", userId)
}
