package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"go-example/restfulapi/config"
	"go-example/restfulapi/database"
	"go-example/restfulapi/handler"
)

func init() {
	log.SetPrefix("[server] ")
	log.SetFlags(log.Lshortfile)

	database.InitializeDatabase()
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/users", handler.UsersHandler).Methods(http.MethodGet)
	r.HandleFunc("/user/{user_id:[0-9]+}", handler.UserForUserIdHandler).Methods(http.MethodGet)
	r.HandleFunc("/user", handler.CreateUserHandler).Methods(http.MethodPost)

	r.HandleFunc("/user/{user_id:[0-9]+}", handler.DeleteUserHandler).Methods(http.MethodDelete)
	r.HandleFunc("/user/{user_id:[0-9]+}", handler.UpdateUserHandler).Methods(http.MethodPut)

	addr := config.GetServerAddr()
	log.Println("start server " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	handler.UsersHandler(w, r)
}
