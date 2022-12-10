package usecases

import (
	"habit/internal/core/domain"
	"habit/internal/core/ports"
	"log"
)

type todoUseCase struct {
	todoRepo ports.TodoRepository
}

func NewTodoUseCase(todoRepo ports.TodoRepository) ports.TodoUseCase {
	return &todoUseCase{todoRepo: todoRepo}
}

func (t *todoUseCase) Get(id int) (*domain.Todo, error) {
	todo, err := t.todoRepo.Get(id)
	if err != nil {
		//TODO Add log here
		return nil, err
	}
	return todo, nil
}

func (t *todoUseCase) GetAll() ([]domain.Todo, error) {
	todos, err := t.todoRepo.GetAll()
	if err != nil {
		//TODO add proper log
		log.Println(err)
		log.Println("in repo error")
		return nil, err
	}
	return todos, err
}

func (t *todoUseCase) Create(title, description string) (*domain.Todo, error) {
	//TODO you should handle the id situation
	todo := domain.NewTodo("1", title, description)

	_, err := t.todoRepo.Create(todo)
	if err != nil {
		//log.Errorw("Error creating from repo", "todo", todo, logging.KeyErr, err)
		return nil, err
	}

	return todo, nil
}
