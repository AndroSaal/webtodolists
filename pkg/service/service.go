package service

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
)

// интерфейс для сервиса для работы с пользователем
type Authorization interface {
	//принимает структуру юзера и возвращает его id в базе и ошибку
	CreateUser(user todo.User) (int, error)

	//получение пользователя по его логину и поролю
	GetUser(username, password string) (todo.User, error)

	//генерация пользователю токена
	GenerateToken(username, password string) (string, error)

	//обработка Токена (парсинг)
	ParseToken(token string) (int, error)
}

// интерфейс для сервиса для работы со списками
type TodoList interface {
	//создание списка по id пользователя, возвращает id
	CreateList(userId int, list todo.TodoList) (int, error)

	//получение всех списков пользователя по его id
	GetAll(userId int) ([]todo.TodoList, error)

	//получение сиска пользователя по id юзера и id списка
	GetById(userId, listId int) (todo.TodoList, error)
}

// интерфейс для работы с item
type TodoItem interface {
}

// Структура с Интерфейсами для общения верхнего слоя с бизнес-логикой
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos),
		TodoList:      NewToDoListService(repos),
	}
}
