package models

import (
	"fmt"
	"sync"
	"time"
)

type Todo struct {
	ID        string `json:"id"`
	Text      string `json:"text"`
	Done      bool   `json:"done"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type TodoStrore struct {
	mu     sync.RWMutex
	todos  map[string]*Todo
	nextID int
}

func NewtodoStore() *TodoStrore {
	return &TodoStrore{
		todos:  make(map[string]*Todo),
		nextID: 1,
	}
}

// GetAll returns all todos
func (s *TodoStrore) GetAll() []*Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		todos = append(todos, todo)
	}

	return todos
}

// GetByID returns a todo by its ID
func (s *TodoStrore) GetByID(id string) *Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.todos[id]
}

// GetByStatus returns todos by their status
func (s *TodoStrore) GetByStatus(done bool) []*Todo {
	s.mu.RLock()
	defer s.mu.RUnlock()

	todos := make([]*Todo, 0, len(s.todos))
	for _, todo := range s.todos {
		if todo.Done == done {
			todos = append(todos, todo)
		}
	}

	return todos
}

// CreateTodo creates a new todo
func (s *TodoStrore) CreateTodo(text string) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	id := fmt.Sprintf("%d", s.nextID)
	s.nextID++

	todo := &Todo{
		ID:        id,
		Text:      text,
		Done:      false,
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}

	s.todos[id] = todo

	return todo
}

// UpdateTodo updates a todo by its ID
func (s *TodoStrore) UpdateTodo(id string, text *string, done *bool) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := s.todos[id]
	if todo == nil {
		return nil
	}

	if text != nil {
		todo.Text = *text
	}

	if done != nil {
		todo.Done = *done
	}

	todo.UpdatedAt = time.Now().Format(time.RFC3339)

	s.todos[id] = todo
	return todo
}

// DeleteTodo deletes a todo by its ID
func (s *TodoStrore) DeleteTodo(id string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.todos[id]; !ok {
		return false
	}

	delete(s.todos, id)
	return true
}

// ToggleTodo toggles a todo by its ID
func (s *TodoStrore) ToggleTodo(id string) *Todo {
	s.mu.Lock()
	defer s.mu.Unlock()

	todo := s.todos[id]
	if todo == nil {
		return nil
	}

	todo.Done = !todo.Done
	todo.UpdatedAt = time.Now().Format(time.RFC3339)

	s.todos[id] = todo
	return todo
}
