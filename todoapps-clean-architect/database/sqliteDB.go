package database

import (
	"log"

	"github.com/jinzhu/gorm"

	"go-example/todoapps-clean-architect/entity"
)

type sqliteDatabase struct{}

func newSqliteDatabase() Database {
	return &sqliteDatabase{}
}

func (s *sqliteDatabase) InitDatabase() {
	db := s.GetDbConnection()
	defer db.Close()
	db.AutoMigrate(&entity.Todos{})
	db.AutoMigrate(&entity.Users{})
}

func (*sqliteDatabase) GetDbConnection() *gorm.DB {
	db, err := gorm.Open("sqlite3", "todos.db")
	if err != nil {
		log.Fatalln(err)
	}
	return db
}
