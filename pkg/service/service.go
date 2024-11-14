package service

import "ToDoApp/pkg/repository"

// интерфейс для работы с пользователем
type Authorization interface {
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
	return &Service{}
}
