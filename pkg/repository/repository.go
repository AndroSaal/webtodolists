package repository

// интерфейс для работы с пользователем
type Authorization interface{
	
}

// интерфейсы для работы со списками
type TodoList interface{

}

// интерфейс для работы с item
type TodoItem interface{

}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
