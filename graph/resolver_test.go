package graph

import (
	"context"
	"math/rand"
	"testing"

	"github.com/TaiwoDevOps/todo-graphql/graph/model"
	"github.com/TaiwoDevOps/todo-graphql/models"
	"github.com/stretchr/testify/assert"
)

func TestCreateTodo(t *testing.T) {

	store := models.NewtodoStore()
	resolver := &Resolver{TodoStrore: store}

	input := model.NewTodo{Text: "Test Todo"}

	todo, err := resolver.Mutation().CreateTodo(context.Background(), input)

	assert.NoError(t, err)
	assert.Equal(t, "Test Todo", todo.Text)
	assert.False(t, todo.Done)
	assert.NotEmpty(t, todo.ID)
}

func TestCreateTodo_Validation_Errors(t *testing.T) {
	store := models.NewtodoStore()

	resolver := &Resolver{TodoStrore: store}

	input := model.NewTodo{Text: ""}
	todo, err := resolver.Mutation().CreateTodo(context.Background(), input)

	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Equal(t, "validation failed for text field: text cannot be empty", err.Error())

	input = model.NewTodo{Text: generateString(256)}
	todo, err = resolver.Mutation().CreateTodo(context.Background(), input)

	assert.Nil(t, todo)
	assert.Error(t, err)
	assert.Equal(t, "validation failed for text field: text cannot be longer than 255 characters", err.Error())

}

func TestGetTodos(t *testing.T) {
	store := models.NewtodoStore()

	resolver := &Resolver{TodoStrore: store}

	store.CreateTodo("Todo 1")
	store.CreateTodo("Todo 2")

	todos, err := resolver.Query().Todos(context.Background())

	assert.NoError(t, err)

	// assert.Equal(t, 2, len(todos))
	assert.Len(t, todos, 2)

}

func generateString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = 'a' + byte(rand.Intn(26))
	}
	return string(b)
}
