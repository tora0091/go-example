package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50); not null"`
	Email   string `gorm:"type:varchar(50); not null"`
	Address string `gorm:"type:varchar(200); not null"`
	Job     string `gorm:"type:varchar(50); not null"`
}

type Users []User
