package model

import (
	"database/sql"
	"fmt"
	"log"

	"go-example/restfulapi/database"
)

const TABLE_NAME = "users"

type User struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Job   string `json:"job"`
}

type Users []User

func NewUser() *User {
	return &User{}
}

func (u *User) CreateUser() error {
	db := database.DBConnection()
	_, err := db.Exec(`insert into `+TABLE_NAME+` (name, email, job) values (?, ?, ?)`, u.Name, u.Email, u.Job)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (u *User) CreateValidator() error {
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

func GetUserList() (Users, error) {
	db := database.DBConnection()
	rows, err := db.Query(`select id, name, email, job from ` + TABLE_NAME + ` where deleted_at is null`)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer rows.Close()

	var users Users
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Job)
		if err != nil {
			log.Println(err)
		}
		users = append(users, user)
	}
	return users, nil
}

func GetUser(userID string) (User, error) {
	db := database.DBConnection()
	row := db.QueryRow(`select id, name, email, job from `+TABLE_NAME+` where id = ? and deleted_at is null`, userID)

	var user User
	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Job)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("data not found")
		}
		return user, err
	}
	return user, nil
}
