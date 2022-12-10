package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"habit/helpers/infra"
	"habit/internal/core/domain"
	"habit/internal/core/ports"
)

type todoMongo struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `bson:"title"`
	Description string             `bson:"description"`
}

func (m *todoMongo) ToDomain() *domain.Todo {

	return &domain.Todo{
		Id:          m.ID.Hex(),
		Title:       m.Title,
		Description: m.Description,
	}
}

func (m todoListMongo) ToDomain() []domain.Todo {
	todos := make([]domain.Todo, len(m))
	for k, td := range m {
		todo := td.ToDomain()
		todos[k] = *todo
	}

	return todos
}

func (m *todoMongo) FromDomain(todo *domain.Todo) {
	if m == nil {
		m = &todoMongo{}
	}

	m.ID = todo.Id
	m.Title = todo.Title
	m.Description = todo.Description
}

type todoListMongo []todoMongo

type todoMongoRepo struct {
	db *mongo.Collection
}

func NewTodoMongoRepo(db *infra.Db) ports.TodoRepository {
	return &todoMongoRepo{db: db.GetCollection("todos")}
}

func (t *todoMongoRepo) Get(id int) (*domain.Todo, error) {
	var todo todoMongo
	result, _ := t.db.Find(context.Background(), bson.M{"_id": id})
	if err := result.Decode(&todo); err != nil {
		return nil, err
	}
	return todo.ToDomain(), nil
}

func (t todoMongoRepo) GetAll() ([]domain.Todo, error) {
	var todos todoListMongo
	result, err := t.db.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	if err := result.All(context.Background(), &todos); err != nil {
		return nil, err
	}

	return todos.ToDomain(), nil
}

func (t todoMongoRepo) Create(todo *domain.Todo) (*domain.Todo, error) {
	var tdMongo *todoMongo = &todoMongo{}
	tdMongo.FromDomain(todo)

	_, err := t.db.InsertOne(context.Background(), tdMongo)

	if err != nil {
		return nil, err
	}

	return todo, nil
}
