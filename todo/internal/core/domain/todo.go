package domain

type Todo struct {
	Id          string
	Title       string
	Description string
}

func NewTodo(id string, title string, description string) *Todo {
	return &Todo{
		Id:          id,
		Title:       title,
		Description: description,
	}
}
