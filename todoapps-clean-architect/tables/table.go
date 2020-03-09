package tables

import "github.com/jinzhu/gorm"

type Todo struct {
	gorm.Model
	Title     string `gorm:"type:varchar(200); not null"`
	Completed bool   `gorm:"default false"`
}

type Todos []Todo

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50); not null"`
	Email   string `gorm:"type:varchar(50); not null"`
	Address string `gorm:"type:varchar(200); not null"`
	Job     string `gorm:"type:varchar(50); not null"`
}

type Users []User
