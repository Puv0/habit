package ports

import "habit/internal/core/domain"

type TodoUseCase interface {
	Get(id int) (*domain.Todo, error)
	GetAll() ([]domain.Todo, error)
	Create(title string, description string) (*domain.Todo, error)
}
