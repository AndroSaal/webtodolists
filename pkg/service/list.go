package service

//сервис для работы со списками

import (
	todo "ToDoApp"
	"ToDoApp/pkg/repository"
)

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
