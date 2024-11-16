package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
)

// интерфейс для работы с пользователем
type Authorization interface {
	//принимает структуру юзера и возвращает его id в базе и ошибку
	CreateUser(user todo.User) (int, error)
}

// интерфейсы для работы со списками
type TodoList interface {
}

// интерфейс для работы с item
type TodoItem interface {
}

// композиция интерфейсов в интерфейсе сервис
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
