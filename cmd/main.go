package main

import (
	"gin-hex/internal/adapters/db"
	"gin-hex/internal/adapters/repository"
	"gin-hex/internal/core/service"
	"gin-hex/internal/handlers"
	"gin-hex/internal/router"
	"log"

	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load() // Load .env variables

    dbInstance := db.ConnectDB() 
    repo := repository.NewTodoRepo(dbInstance)
    todoService := service.NewTodoService(repo)
    todoHandler := handlers.NewTodoHandler(todoService)

    r := router.SetupRouter(todoHandler)

    log.Println("Server running on port 3000")
    r.Run(":3000") // Start Gin server
}
