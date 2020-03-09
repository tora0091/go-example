package repository

import (
	"fmt"
	"go-example/todoapps-clean-architect/database"
	"go-example/todoapps-clean-architect/entity"
)

type UsersRepository interface {
	Save(user *entity.User) (*entity.User, error)
	FindAll() (*entity.Users, error)
}

type usersRepository struct{}

func NewUsersRepository() UsersRepository {
	return &usersRepository{}
}

func (*usersRepository) Save(user *entity.User) (*entity.User, error) {
	db := database.GetDbConnection()
	db.NewRecord(&user)
	db.Create(user)
	if db.NewRecord(&user) == false {
		return user, nil
	}
	return nil, fmt.Errorf("failed create user")
}

func (*usersRepository) FindAll() (*entity.Users, error) {
	db := database.GetDbConnection()
	defer db.Close()

	var users entity.Users
	db.Find(&users)

	return &users, nil
}
