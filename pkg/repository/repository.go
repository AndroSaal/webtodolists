package repository

import (
	todo "ToDoApp/entities"

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
	GetAllList(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	DeleteById(userId, listId int) error
	UpdateById(userId, listId int, list todo.UpdateListInput) error
}

// интерфейс для работы с item
type TodoItem interface {
	Create(listId int, item todo.TodoItem) (int, error)
	GetAll(listId int) ([]todo.TodoItem, error)
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
		TodoItem:      NewItemPostgres(db),
	}

}
