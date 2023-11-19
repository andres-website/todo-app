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
	}
}
