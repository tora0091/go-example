package repository

import (
	"fmt"

	"go-example/todoapps-clean-architect/database"
	"go-example/todoapps-clean-architect/entity"
)

type TodosRepository interface {
	Save(todo *entity.Todo) (*entity.Todo, error)
	FindAll() (*entity.Todos, error)
	FindByID(id int) (*entity.Todo, error)
	UpdateByID(id int, updateData *entity.Todo) (*entity.Todo, error)
	DeleteByID(id int) (*entity.Todo, error)
}

type todosRepository struct {
	dbDespatcher database.Database
}

func NewTodosRepository() TodosRepository {
	return &todosRepository{
		dbDespatcher: database.NewDatabase(),
	}
}

func (r *todosRepository) Save(todo *entity.Todo) (*entity.Todo, error) {
	db := r.dbDespatcher.GetDbConnection()
	db.NewRecord(todo)
	db.Create(&todo)
	if db.NewRecord(todo) == false {
		return todo, nil
	}
	return nil, fmt.Errorf("failed create todo")
}

func (r *todosRepository) FindAll() (*entity.Todos, error) {
	db := r.dbDespatcher.GetDbConnection()
	defer db.Close()

	var todos entity.Todos
	db.Find(&todos)

	return &todos, nil
}

func (r *todosRepository) FindByID(id int) (*entity.Todo, error) {
	db := r.dbDespatcher.GetDbConnection()
	defer db.Close()

	var todo entity.Todo
	recordNotFound := db.First(&todo, id).RecordNotFound()
	if recordNotFound {
		return nil, fmt.Errorf("record not found")
	}
	return &todo, nil
}

func (r *todosRepository) UpdateByID(id int, updateData *entity.Todo) (*entity.Todo, error) {
	var todo entity.Todo
	db := r.dbDespatcher.GetDbConnection()

	if err := db.First(&todo, id).Update(updateData).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

func (r *todosRepository) DeleteByID(id int) (*entity.Todo, error) {
	var todo entity.Todo
	db := r.dbDespatcher.GetDbConnection()
	if err := db.Delete(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
