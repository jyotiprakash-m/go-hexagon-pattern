package router

import (
	"gin-hex/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(todoHandler *handlers.TodoHandler) *gin.Engine {
    r := gin.Default()

    r.POST("/todos", todoHandler.CreateTodo)
    r.GET("/todos", todoHandler.GetTodos)
    r.GET("/todos/:id", todoHandler.GetTodoByID)
    r.DELETE("/todos/:id", todoHandler.DeleteTodo)
    r.PUT("/todos/:id", todoHandler.UpdateTodo)

    return r
}
