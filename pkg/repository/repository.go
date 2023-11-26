package repository

import (
	todo "github.com/andres-website/todo-app/pkg"
	"github.com/jmoiron/sqlx"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
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
		TodoList:     NewTodoListPostgres(db),
	}
}
