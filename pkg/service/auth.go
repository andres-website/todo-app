package service

import (
	"crypto/sha1"
	"fmt"

	todo "github.com/andres-website/todo-app/pkg"
	"github.com/andres-website/todo-app/pkg/repository"
)

const salt = "lkjl3lkjpjh3kjbh3lkj3lkje"

type AuthService struct {
	repo repository.Autorization
}

func NewAuthService(repo repository.Autorization) *AuthService {

	return &AuthService{repo: repo}
}

// todo - означает что package todo (в файле pkg/user.go)
func (s *AuthService) CreateUser(user todo.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {

	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprint("%x", hash.Sum([]byte(salt)))
}
