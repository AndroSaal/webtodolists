package service

import (
	todo "ToDoApp/entities"
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
	GetAllList(userId int) ([]todo.TodoList, error)

	//получение сиска пользователя по id юзера и id списка
	GetById(userId, listId int) (todo.TodoList, error)

	//Удаление списка по id
	DeleteById(userId, ListId int) error

	UpdateById(userId, listId int, input todo.UpdateListInput) error
}

// интерфейс для работы с item
type TodoItem interface {
	//создание задачи, возвращает id созданной задачи
	Create(userId, listId int, item todo.TodoItem) (int, error)

	//получение всех задач из списка пользователя
	GetAll(userId, listId int) ([]todo.TodoItem, error)
}

// Структура с Интерфейсами для общения верхнего слоя с бизнес-логикой
type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewToDoListService(repos.TodoList),
		TodoItem:      NewToDoItemService(repos.TodoItem, repos.TodoList),
	}
}
