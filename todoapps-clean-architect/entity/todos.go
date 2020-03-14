package entity

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"type:varchar(200); not null" json:"title"`
	Completed bool   `gorm:"default false" json:"completed"`
}

type Todos []Todo
