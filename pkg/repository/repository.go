package repository

import (
	todo "github.com/andres-website/todo-app/pkg"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
}

type TodoList interface {
}

type TodoItems interface {
}

type Repository struct {
	Autorization
	TodoList
	TodoItems
}

func NewRepository(db *sqlx.DB) *Repository {

	return &Repository{
		Autorization: NewAuthPostgres(db),
	}
}
