package repository

import (
	"tdp-api/internal/model"

	"gorm.io/gorm"
)

// TaskRepository интерфейс репозитория задач
type TaskRepository interface {
	GetTasks(userId uint) ([]model.Task, error)
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

func (r *taskRepository) GetTasks(userId uint) ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Preload("TaskCategory").Where("user_id = ?", userId).Find(&tasks)
	return tasks, result.Error
}

func (r *taskRepository) GetTask(id uint) (model.Task, error) {
	var task model.Task
	result := r.db.Preload("TaskCategory").First(&task, id)
	return task, result.Error
}

func (r *taskRepository) CreateTask(task model.Task) (model.Task, error) {
	result := r.db.Preload("TaskCategory").Create(&task)
	result = r.db.Preload("TaskCategory").First(&task, task.ID)
	return task, result.Error
}

func (r *taskRepository) UpdateTask(task model.Task) (model.Task, error) {
	result := r.db.Preload("TaskCategory").Save(&task)
	result = r.db.Preload("TaskCategory").First(&task, task.ID)
	return task, result.Error
}

func (r *taskRepository) DeleteTask(id uint) error {
	result := r.db.Preload("TaskCategory").Delete(&model.Task{}, id)
	return result.Error
}

// TemplateTaskRepository интерфейс репозитория задач
type TemplateTaskRepository interface {
	GetTasks(userId uint) ([]model.TemplateTask, error)
	GetTask(id uint) (model.TemplateTask, error)
	CreateTask(task model.TemplateTask) (model.TemplateTask, error)
	UpdateTask(task model.TemplateTask) (model.TemplateTask, error)
	DeleteTask(id uint) error
}

type templateTaskRepository struct {
	db *gorm.DB
}

// NewTemplateTaskRepository создает новый экземпляр репозитория задач
func NewTemplateTaskRepository(db *gorm.DB) TemplateTaskRepository {
	return &templateTaskRepository{db: db}
}

func (r *templateTaskRepository) GetTasks(userId uint) ([]model.TemplateTask, error) {
	var tasks []model.TemplateTask
	result := r.db.Preload("TaskCategory").Where("user_id = ?", userId).Find(&tasks)
	return tasks, result.Error
}

func (r *templateTaskRepository) GetTask(id uint) (model.TemplateTask, error) {
	var task model.TemplateTask
	result := r.db.Preload("TaskCategory").First(&task, id)
	return task, result.Error
}

func (r *templateTaskRepository) CreateTask(task model.TemplateTask) (model.TemplateTask, error) {
	result := r.db.Preload("TaskCategory").Create(&task)
	result = r.db.Preload("TaskCategory").First(&task, task.ID)
	return task, result.Error
}

func (r *templateTaskRepository) UpdateTask(task model.TemplateTask) (model.TemplateTask, error) {
	result := r.db.Preload("TaskCategory").Save(&task)
	result = r.db.Preload("TaskCategory").First(&task, task.ID)
	return task, result.Error
}

func (r *templateTaskRepository) DeleteTask(id uint) error {
	result := r.db.Delete(&model.Task{}, id)
	return result.Error
}
