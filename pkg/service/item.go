package service

import (
	todo "ToDoApp/entities"
	"ToDoApp/pkg/repository"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewToDoItemService(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{
		repo:     repo,
		listRepo: listRepo,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {

	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		//list does not even exist or not belong to user
		return 0, err
	}

	//передаём на уровень ниже - в уровень репозитория
	return s.repo.Create(listId, item)
}
