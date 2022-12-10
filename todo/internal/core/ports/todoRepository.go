package ports

import "habit/internal/core/domain"

type TodoRepository interface {
	Get(id int) (*domain.Todo, error)
	GetAll() ([]domain.Todo, error)
	Create(todo *domain.Todo) (*domain.Todo, error)
}
