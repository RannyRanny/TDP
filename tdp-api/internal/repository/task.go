package repository

import (
	"tdp-api/internal/model"

	"gorm.io/gorm"
)

// TaskRepository интерфейс репозитория задач
type TaskRepository interface {
	GetTasks() ([]model.Task, error)
	GetTask(id uint) (model.Task, error)
	CreateTask(task model.Task) (model.Task, error)
	UpdateTask(task model.Task) (model.Task, error)
	DeleteTask(id uint) error
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository создает новый экземпляр репозитория задач
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetTasks() ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Find(&tasks)
	r.db.Preload("TaskCategory").Find(&tasks)
	return tasks, result.Error
}

func (r *taskRepository) GetTask(id uint) (model.Task, error) {
	var task model.Task
	r.db.Preload("TaskCategory").Find(&task)
	result := r.db.First(&task, id)
	return task, result.Error
}

func (r *taskRepository) CreateTask(task model.Task) (model.Task, error) {
	result := r.db.Create(&task)
	return task, result.Error
}

func (r *taskRepository) UpdateTask(task model.Task) (model.Task, error) {
	result := r.db.Save(&task)
	return task, result.Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	result := r.db.Delete(&model.Task{}, id)
	return result.Error
}
