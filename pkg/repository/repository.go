package repository

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

func NewRepository() *Repository {

	return &Repository{}
}
