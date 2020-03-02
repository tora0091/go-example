package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"go-example/restfulapi/auth"
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

	// jwt auth key
	r.HandleFunc("/auth", auth.GetAuthTokenHandler)

	api := r.PathPrefix("/api").Subrouter()
	// csrfMiddleware := csrf.Protect([]byte("32-byte-long-auth-key"))
	api.Use(auth.AuthMiddleware)
	// api.Use(csrfMiddleware)
	api.HandleFunc("/users", handler.UsersHandler).Methods(http.MethodGet)
	api.HandleFunc("/user/{user_id:[0-9]+}", handler.UserForUserIdHandler).Methods(http.MethodGet)
	api.HandleFunc("/user", handler.CreateUserHandler).Methods(http.MethodPost)

	api.HandleFunc("/user/{user_id:[0-9]+}", handler.DeleteUserHandler).Methods(http.MethodDelete)
	api.HandleFunc("/user/{user_id:[0-9]+}", handler.UpdateUserHandler).Methods(http.MethodPut)

	addr := config.GetServerAddr()
	log.Println("start server " + addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
