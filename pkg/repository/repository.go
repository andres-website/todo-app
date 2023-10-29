package repository

import "github.com/jmoiron/sqlx"

type Autorization interface {
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

	return &Repository{}
}
