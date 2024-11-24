// имплементация логики сервиса по созданиб задач на уровне репозитория
package repository

import (
	todo "ToDoApp/entities"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (t *TodoItemPostgres) Create(listId, item todo.TodoItem) (int, error) {
	tx, err := t.db.Begin() //старутем транзакцию
	if err != nil {
		return 0, err
	}

	var itemId int
	//формируем текст запроса в БД, занести в табл с задачами заголовок и описание заддачи
	createItemQuery := fmt.Sprintf(`INSERT INTO %s (title, description) VALUES ($1, $2) RETURNING id`, todoItemsTable)
	//запрос в БД
	row := tx.QueryRow(createItemQuery, item.Title, item.Description)
	//получение данных из БД в itemId
	err = row.Scan(&itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// формируем запрос в БД, обозначаем связь конкретной задачи с конкретнымы списком
	createListItemQuery := fmt.Sprintf(`INSERT INTO %s (list_id, item_id) VALUES ($1, $2)`, listsItemsTable)
	_, err = tx.Exec(createListItemQuery, listId, itemId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, tx.Commit()
}
