package repository

import (
	todo "ToDoApp"
	"fmt"

	"github.com/jmoiron/sqlx"
)

//имплементация интерфейса ToDoList

type ToDoListPostgres struct {
	db *sqlx.DB
}

func NewListPostgres(db *sqlx.DB) *ToDoListPostgres {
	return &ToDoListPostgres{db: db}
}

func (t *ToDoListPostgres) CreateList(userId int, list todo.TodoList) (int, error) {
	//инициализируем транзакцию
	tx, err := t.db.Begin()

	if err != nil {
		return 0, err
	}

	var id int

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id", todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	//если возникают какие-то ошибки - возвращаем все назад
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES ($1, $2)", usersListsTable)

	//вополнение запроса, без чтение возвращаемой информации, если возникают ошибки
	//то откатываем транзакцию
	if _, err := tx.Exec(createUsersListQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}
