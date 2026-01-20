package main

import (
	"github.com/gin-gonic/gin"
	"github.com/higorworkout/todo-api/internal/adapters/inbound/http"
	dbAdapter "github.com/higorworkout/todo-api/internal/adapters/outbound/database"
	"github.com/higorworkout/todo-api/internal/application"
	"github.com/higorworkout/todo-api/internal/infrastructure/database"
)


func main() {
	db := database.NewPostgres()
	repo := dbAdapter.NewTaskGormRepository(db)
	uc := application.NewTaskUseCase(repo)

	r := gin.Default()
	http.NewTaskHandler(r, uc)

	r.Run(":8080")
}