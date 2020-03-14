package entity

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name    string `gorm:"type:varchar(50); not null" json:"name"`
	Email   string `gorm:"type:varchar(50); not null" json:"email"`
	Address string `gorm:"type:varchar(200); not null" json:"address"`
	Job     string `gorm:"type:varchar(50); not null" json:"job"`
}

type Users []User
