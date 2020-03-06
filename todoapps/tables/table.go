package tables

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"type:varchar(200); not null"`
	Completed bool   `gorm:"default false"`
}

type Todos []Todo
