package model

import (
	"fmt"
	"log"

	"go-example/restfulapi/database"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Job   string `json:"job"`
}

type Users []User

func (u User) CreateUser() error {
	db := database.DBConnection()
	_, err := db.Exec(`insert into users (name, email, job) values (?, ?, ?)`, u.Name, u.Email, u.Job)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (u User) CreateValidator() error {
	if u.Name == "" {
		return fmt.Errorf("name is required")
	}
	if u.Email == "" {
		return fmt.Errorf("email is required")
	}
	if u.Job == "" {
		return fmt.Errorf("job is required")
	}
	return nil
}
