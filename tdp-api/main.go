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

	db.AutoMigrate(&model.Task{})
	db.AutoMigrate(&model.TaskCategory{})
	db.AutoMigrate(&model.TemplateDay{})
	err = db.AutoMigrate(&model.TelegramAuthData{})

	r := gin.Default()

	taskRepository := repository.NewTaskRepository(db)
	taskHandler := handler.NewTaskHandler(taskRepository)

	r.GET("/tasks", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	templateTaskRepository := repository.NewTemplateTaskRepository(db)
	templateTaskHandler := handler.NewTemplateTaskHandler(templateTaskRepository)

	r.GET("/tasks/template", templateTaskHandler.GetTasks)
	r.GET("/tasks/template/:id", templateTaskHandler.GetTask)
	r.POST("/tasks/template", templateTaskHandler.CreateTask)
	r.PUT("/tasks/template/:id", templateTaskHandler.UpdateTask)
	r.DELETE("/tasks/template/:id", templateTaskHandler.DeleteTask)

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

	templateDayRepository := repository.NewGormTemplateDayRepository(db)
	templateDayHandler := handler.NewTemplateDayHandler(templateDayRepository)
	r.POST("/day/template", templateDayHandler.CreateTemplateDay)
	r.GET("/day/template/:id", templateDayHandler.GetTemplateDay)
	r.PUT("/day/template/:id", templateDayHandler.UpdateTemplateDay)
	r.DELETE("/day/template/:id", templateDayHandler.DeleteTemplateDay)

	r.Run()
}
