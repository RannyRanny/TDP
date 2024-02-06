package main

import (
	"tdp-api/internal/config"
	"tdp-api/internal/handler"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files" // swagger embed files
	"github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	_ "tdp-api/docs"
)

func main() {
	cfg := config.New()
	db, err := gorm.Open(postgres.Open(cfg.DSN), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&model.Task{}, &model.TaskCategory{}, &model.TelegramAuthData{})

	r := gin.Default()

	taskRepository := repository.NewTaskRepository(db)
	taskHandler := handler.NewTaskHandler(taskRepository)

	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	taskCategoryRepository := repository.NewTaskCategoryRepository(db)
	taskCategoryHandler := handler.NewTaskCategoryHandler(taskCategoryRepository)

	r.GET("/categories", taskCategoryHandler.GetCategories)
	r.GET("/categories/:id", taskCategoryHandler.GetCategory)
	r.POST("/categories", taskCategoryHandler.CreateCategory)
	r.PUT("/categories/:id", taskCategoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", taskCategoryHandler.DeleteCategory)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	telegramRepository := repository.NewGormTelegramAuthRepository(db)
	telegramAuthHandler := handler.NewTelegramHandler(telegramRepository)
	r.POST("/telegram_auth", telegramAuthHandler.TelegramAuth)

	r.Run() // listen and serve on 0.0.0.0:8080
}
