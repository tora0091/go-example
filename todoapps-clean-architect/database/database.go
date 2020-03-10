package database

import (
	"github.com/jinzhu/gorm"
)

type Database interface {
	InitDatabase()
	GetDbConnection() *gorm.DB
}

func NewDatabase() Database {
	return newSqliteDatabase()
}
