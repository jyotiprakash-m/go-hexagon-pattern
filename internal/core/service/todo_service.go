package service

import (
	"gin-hex/internal/core/models"
	"gin-hex/internal/core/ports"
)

type TodoService struct {
    Repo ports.TodoRepository
}

func NewTodoService(repo ports.TodoRepository) *TodoService {
    return &TodoService{Repo: repo}
}

func (s *TodoService) CreateTodo(title string) (*models.Todo, error) {
    todo := &models.Todo{Title: title}
    err := s.Repo.Create(todo)
    return todo, err
}

func (s *TodoService) GetTodos() ([]models.Todo, error) {
    return s.Repo.GetAll()
}

func (s *TodoService) GetTodoByID(id uint) (*models.Todo, error) {
    return s.Repo.GetByID(id)
}

// DeleteTodo
func (s *TodoService) DeleteTodo(id uint) error {
    return s.Repo.DeleteTodo(id)
}
//UpdateTodo
func (s *TodoService) UpdateTodo(todo *models.Todo) error {
    return s.Repo.UpdateTodo(todo)
}