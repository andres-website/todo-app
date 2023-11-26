package service

import (
	todo "github.com/andres-website/todo-app/pkg"
	"github.com/andres-website/todo-app/pkg/repository"
)

type Autorization interface {
	CreateUser(user todo.User) (int, error)
	GentrateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
}

type TodoItem interface {
}

type Service struct {
	Autorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {

	return &Service{
		Autorization: NewAuthService(repos.Autorization),
		TodoList:     NewTodoListService(repos.TodoList),
	}
}
