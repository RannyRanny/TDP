package main

import (
	"tdp-api/internal/config"
	"tdp-api/internal/handler"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"
	service "tdp-api/internal/services"

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

	db.AutoMigrate(&model.Task{}, &model.TemplateTask{})
	db.AutoMigrate(&model.TaskCategory{})
	db.AutoMigrate(&model.TemplateDay{}, &model.Day{})
	db.AutoMigrate(&model.TelegramAuthData{})

	r := gin.Default()

	taskRepository := repository.NewTaskRepository(db)
	taskHandler := handler.NewTaskHandler(taskRepository)

	r.GET("/tasks/user/:userId", taskHandler.GetTasks)
	r.GET("/tasks/:id", taskHandler.GetTask)
	r.POST("/tasks", taskHandler.CreateTask)
	r.PUT("/tasks/:id", taskHandler.UpdateTask)
	r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	templateTaskRepository := repository.NewTemplateTaskRepository(db)
	templateTaskHandler := handler.NewTemplateTaskHandler(templateTaskRepository)

	r.GET("/tasks/template/user/:userId", templateTaskHandler.GetTasks)
	r.GET("/tasks/template/:id", templateTaskHandler.GetTask)
	r.POST("/tasks/template", templateTaskHandler.CreateTask)
	r.PUT("/tasks/template/:id", templateTaskHandler.UpdateTask)
	r.DELETE("/tasks/template/:id", templateTaskHandler.DeleteTask)

	templateDayRepository := repository.NewGormTemplateDayRepository(db)
	templateDayHandler := handler.NewTemplateDayHandler(templateDayRepository)
	r.GET("/day/template/user/:userId", templateDayHandler.GetTemplateDays)
	r.GET("/day/template/:id", templateDayHandler.GetTemplateDay)
	r.POST("/day/template", templateDayHandler.CreateTemplateDay)
	r.PUT("/day/template/:id", templateDayHandler.UpdateTemplateDay)
	r.DELETE("/day/template/:id", templateDayHandler.DeleteTemplateDay)

	dayRepository := repository.NewGormDayRepository(db)
	dayHandler := handler.NewDayHandler(dayRepository)
	r.GET("/day/user/:userId", dayHandler.GetDays)
	r.GET("/day/user/:userId/:date", dayHandler.GetDaysByDate)
	r.POST("/day", dayHandler.CreateDay)
	r.GET("/day/:id", dayHandler.GetDay)
	r.PUT("/day/:id", dayHandler.UpdateDay)
	r.DELETE("/day/:id", dayHandler.DeleteDay)

	taskCategoryRepository := repository.NewTaskCategoryRepository(db)
	taskCategoryHandler := handler.NewTaskCategoryHandler(taskCategoryRepository)

	r.GET("/categories/user/:userId", taskCategoryHandler.GetCategories)
	r.GET("/categories/:id", taskCategoryHandler.GetCategory)
	r.POST("/categories", taskCategoryHandler.CreateCategory)
	r.PUT("/categories/:id", taskCategoryHandler.UpdateCategory)
	r.DELETE("/categories/:id", taskCategoryHandler.DeleteCategory)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	telegramRepository := repository.NewGormTelegramAuthRepository(db)
	telegramService := service.NewTelegramService(telegramRepository)
	telegramAuthHandler := handler.NewTelegramHandler(telegramService)
	r.POST("/telegram/auth", telegramAuthHandler.TelegramAuth)
	r.POST("/user/create", telegramAuthHandler.CreateUser)

	r.Run()
}
