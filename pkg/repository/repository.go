package repository

import (
	todo "ToDoApp"

	"github.com/jmoiron/sqlx"
)

// интерфейс для работы с пользователем
type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

// интерфейсы для работы со списками
type TodoList interface {
	CreateList(userId int, list todo.TodoList) (int, error)
}

// интерфейс для работы с item
type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewListPostgres(db),
	}

}
