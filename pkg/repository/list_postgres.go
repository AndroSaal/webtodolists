//имплементация интерфейса ToDoList из repository.go

package repository

import (
	todo "ToDoApp/entities"
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
)

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

	createListQuery := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`, todoListTable)
	row := tx.QueryRow(createListQuery, list.Title, list.Description)

	//если возникают какие-то ошибки - возвращаем все назад
	if err := row.Scan(&id); err != nil {
		tx.Rollback()
		return 0, err
	}

	createUsersListQuery := fmt.Sprintf(`INSERT INTO %s (user_id, list_id) VALUES ($1, $2)`, usersListsTable)

	//вополнение запроса, без чтение возвращаемой информации, если возникают ошибки
	//то откатываем транзакцию
	if _, err := tx.Exec(createUsersListQuery, userId, id); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, tx.Commit()
}

func (t *ToDoListPostgres) GetAllList(userId int) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf(`SELECT tdlst.id, tdlst.title, tdlst.description FROM %s tdlst 
							INNER JOIN %s usrlst ON tdlst.id = usrlst.list_id WHERE usrlst.user_id = $1`,
		todoListTable, usersListsTable)
	err := t.db.Select(&lists, query, userId)

	return lists, err
}

func (t *ToDoListPostgres) GetById(userId, listId int) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`SELECT tdlst.id, tdlst.title, tdlst.description FROM %s tdlst 
							INNER JOIN %s usrlst ON tdlst.id = usrlst.list_id 
							WHERE usrlst.user_id = $1 AND tdlst.id = $2`,
		todoListTable, usersListsTable)

	err := t.db.Get(&list, query, userId, listId)

	return list, err

}

func (t *ToDoListPostgres) DeleteById(userId, listId int) error {
	query := fmt.Sprintf(`DELETE FROM %s tl USING %s ul 
							WHERE tl.id = ul.list_id AND ul.user_id = $1 AND ul.list_id = $2`,
		todoListTable, usersListsTable)

	_, err := t.db.Exec(query, userId, listId)

	return err
}

func (t *ToDoListPostgres) UpdateById(userId, listId int, list todo.UpdateListInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if list.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, *list.Title)
		argId++
	}

	if list.Description != nil {
		setValues = append(setValues, fmt.Sprintf("description=$%d", argId))
		args = append(args, *list.Description)
		argId++
	}

	setQuery := strings.Join(setValues, ",")

	query := fmt.Sprintf(`UPDATE %s tl SET %s FROM %s ul WHERE tl.id = ul.list_id AND ul.list_id = $%d AND ul.user_id=$%d`, todoListTable, setQuery, usersListsTable, argId, argId+1)

	args = append(args, listId, userId)
	_, err := t.db.Exec(query, args...)

	return err

}
