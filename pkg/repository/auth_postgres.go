package repository

import (
	todo "ToDoApp"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// тип реализующий интерфейс
type AuthPostgres struct {
	db *sqlx.DB
}

// конструктор этого типа
func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{
		db: db,
	}
}

// метод CreateUser, чтобы AuthPostgres реализовывал интерфейс
func (r *AuthPostgres) CreateUser(user todo.User) (int, error) {
	var id int

	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) VALUES ($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return 0, nil
}
