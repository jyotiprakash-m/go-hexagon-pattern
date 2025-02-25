package repository

import (
	"errors"
	"gin-hex/internal/core/models"

	"gorm.io/gorm"
)

type TodoRepo struct {
    DB *gorm.DB
}

func NewTodoRepo(db *gorm.DB) *TodoRepo {
    return &TodoRepo{DB: db}
}

func (r *TodoRepo) Create(todo *models.Todo) error {
    return r.DB.Create(todo).Error
}

func (r *TodoRepo) GetAll() ([]models.Todo, error) {
    var todos []models.Todo
    err := r.DB.Find(&todos).Error
    return todos, err
}
func (r *TodoRepo) GetByID(id uint) (*models.Todo, error) {
    var todo models.Todo
    result := r.DB.First(&todo, id)
    if result.Error != nil {
        if errors.Is(result.Error, gorm.ErrRecordNotFound) {
            return nil, errors.New("todo not found")
        }
        return nil, result.Error
    }
    return &todo, nil
}
// DeleteTodo
func (r *TodoRepo) DeleteTodo(id uint) error {
    result := r.DB.Delete(&models.Todo{}, id)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("todo not found")
    }
    return nil
}
// UpdateTodo
func (r *TodoRepo) UpdateTodo(todo *models.Todo) error {
    result := r.DB.Save(todo)
    if result.Error != nil {
        return result.Error
    }
    if result.RowsAffected == 0 {
        return errors.New("todo not found")
    }
    return nil
}