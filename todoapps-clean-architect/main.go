package main

import (
	"net/http"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"go-example/todoapps-clean-architect/database"
	"go-example/todoapps-clean-architect/router"
)

var (
	dbDispater database.Database = database.NewDatabase()
)

func init() {
	dbDispater.InitDatabase()
}

func main() {
	r := router.InitRouter()
	s := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
	s.ListenAndServe()
}
