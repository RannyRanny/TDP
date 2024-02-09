package repository

import (
	"gorm.io/gorm"
)

// ITaskRepository интерфейс репозитория задач
type ITaskRepository[T any] interface {
	GetTasks(userId uint) ([]T, error)
	GetTask(id uint) (T, error)
	CreateTask(task T) (T, error)
	UpdateTask(task T) (T, error)
	DeleteTask(id uint) error
}

type TaskRepository[T any] struct {
	db *gorm.DB
}

// NewTaskRepository создает новый экземпляр репозитория задач
func NewTaskRepository[T any](db *gorm.DB) *TaskRepository[T] {
	return &TaskRepository[T]{db: db}
}

func (r *TaskRepository[T]) GetTasks(userId uint) ([]T, error) {
	var tasks []T
	result := r.db.Preload("TaskCategory").Where("user_id = ?", userId).Find(&tasks)
	return tasks, result.Error
}

func (r *TaskRepository[T]) GetTask(id uint) (T, error) {
	var task T
	result := r.db.Preload("TaskCategory").First(&task, id)
	return task, result.Error
}

func (r *TaskRepository[T]) CreateTask(task T) (T, error) {
	result := r.db.Preload("TaskCategory").Create(&task)
	return task, result.Error
}

func (r *TaskRepository[T]) UpdateTask(task T) (T, error) {
	result := r.db.Preload("TaskCategory").Save(&task)
	return task, result.Error
}

func (r *TaskRepository[T]) DeleteTask(id uint) error {
	var task T
	_ = r.db.Preload("TaskCategory").First(&task, id)
	result := r.db.Preload("TaskCategory").Delete(task, id)
	return result.Error
}
