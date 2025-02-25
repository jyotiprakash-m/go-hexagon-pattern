package handlers

import (
	"gin-hex/internal/core/models"
	"gin-hex/internal/core/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
    Service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
    return &TodoHandler{Service: s}
}

// Create a new Todo
func (h *TodoHandler) CreateTodo(c *gin.Context) {
    var req struct{ Title string }
    
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

    todo, err := h.Service.CreateTodo(req.Title)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(201, todo)
}

// Get all Todos
func (h *TodoHandler) GetTodos(c *gin.Context) {
    todos, err := h.Service.GetTodos()
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, todos)
}

// ✅ New Handler to Get Todo by ID
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(500, gin.H{"error": "Invalid ID"})
        return
    }

    todo, err := h.Service.GetTodoByID(uint(id))
    if err != nil {
        c.JSON(404, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, todo)
}
// DeleteTodo
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
    idParam := c.Param("id")
    id, err := strconv.Atoi(idParam)
    if err != nil {
        c.JSON(500, gin.H{"error": "Invalid ID"})
        return
    }

    err = h.Service.DeleteTodo(uint(id))
    if err != nil {
        c.JSON(404, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"message": "Todo deleted successfully"})
}
// UpdateTodo
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
    var req struct {
        ID    uint   `json:"id"`
        Title string `json:"title"`
    }
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(400, gin.H{"error": "Invalid request"})
        return
    }

    todo := &models.Todo{ID: req.ID, Title: req.Title}
    err := h.Service.UpdateTodo(todo)
    if err != nil {
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    c.JSON(200, todo)
}