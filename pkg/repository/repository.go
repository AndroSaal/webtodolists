package repository

// интерфейс для работы с пользователем
type Authorization interface{}

// интерфейсы для работы со списками
type TodoList interface{}

// интерфейс для работы с item
type TodoItem interface{}

//композиция интерфейсов в интерфейсе сервис
// type Repository interface {
// 	Authorization
// 	TodoList
// 	TodoItem
// }

// type someRepository struct {
// 	someField string
// }

// func NewSomeRepository() *Repository {
// 	return &someRepository{
// 		someField: "someField",

// 	}

// 	}
// }

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository() *Repository {
	return &Repository{}
}
