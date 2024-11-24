package service

//сервис для работы со списками

import (
	todo "ToDoApp/entities"
	"ToDoApp/pkg/repository"
)

// имплементация интерфейса
type TodoListService struct {
	repo repository.TodoList
}

func NewToDoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{
		repo: repo,
	}
}

// передаем на нижний уровень
func (t *TodoListService) CreateList(userId int, list todo.TodoList) (int, error) {
	return t.repo.CreateList(userId, list)
}

func (t *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return t.repo.GetAll(userId)
}

func (t *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return t.repo.GetById(userId, listId)
}

func (t *TodoListService) DeleteById(userId, listId int) error {
	return t.repo.DeleteById(userId, listId)
}

func (t *TodoListService) UpdateById(userId, listId int, list todo.UpdateListInput) error {
	return t.repo.UpdateById(userId, listId, list)
}
