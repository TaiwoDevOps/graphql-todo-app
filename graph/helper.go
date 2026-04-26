package graph

import (
	"github.com/TaiwoDevOps/todo-graphql/graph/model"
	"github.com/TaiwoDevOps/todo-graphql/models"
)

func convertTodo(todo *models.Todo) *model.Todo {
	if todo == nil {
		return nil
	}

	return &model.Todo{
		ID:        todo.ID,
		Text:      todo.Text,
		Done:      todo.Done,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}
}

func convertTodos(todos []*models.Todo) []*model.Todo {
	if todos == nil {
		return nil
	}

	converted := make([]*model.Todo, len(todos))
	for i, todo := range todos {
		converted[i] = convertTodo(todo)
	}

	return converted
}
