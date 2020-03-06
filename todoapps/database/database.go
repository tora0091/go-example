package database

import (
	"log"

	"github.com/jinzhu/gorm"

	"go-example/todoapps/tables"
)

func InitDatabase() {
	db := GetDbConnection()
	defer db.Close()
	db.AutoMigrate(&tables.Todos{})
}

func GetDbConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
