package ports

import "gin-hex/internal/core/models"

type TodoRepository interface {
    Create(todo *models.Todo) error
    GetAll() ([]models.Todo, error)
    GetByID(id uint) (*models.Todo, error)
    DeleteTodo(id uint) error
    UpdateTodo(todo *models.Todo) error
}
