package repository

import (
	"tdp-api/internal/model"

	"gorm.io/gorm"
)

// TaskCategoryRepository интерфейс репозитория категорий задач
type TaskCategoryRepository interface {
	GetCategories(userId uint) ([]model.TaskCategory, error)
	GetCategory(id uint) (model.TaskCategory, error)
	CreateCategory(category model.TaskCategory) (model.TaskCategory, error)
	UpdateCategory(category model.TaskCategory) (model.TaskCategory, error)
	DeleteCategory(id uint) error
}

type taskCategoryRepository struct {
	db *gorm.DB
}

// NewTaskCategoryRepository создает новый экземпляр репозитория категорий задач
func NewTaskCategoryRepository(db *gorm.DB) TaskCategoryRepository {
	return &taskCategoryRepository{db: db}
}

func (r *taskCategoryRepository) GetCategories(userId uint) ([]model.TaskCategory, error) {
	var categories []model.TaskCategory
	result := r.db.Where("user_id = ?", userId).Find(&categories)
	return categories, result.Error
}

func (r *taskCategoryRepository) GetCategory(id uint) (model.TaskCategory, error) {
	var category model.TaskCategory
	result := r.db.First(&category, id)
	return category, result.Error
}

func (r *taskCategoryRepository) CreateCategory(category model.TaskCategory) (model.TaskCategory, error) {
	result := r.db.Create(&category)
	return category, result.Error
}

func (r *taskCategoryRepository) UpdateCategory(category model.TaskCategory) (model.TaskCategory, error) {
	result := r.db.Save(&category)
	return category, result.Error
}

func (r *taskCategoryRepository) DeleteCategory(id uint) error {
	result := r.db.Delete(&model.TaskCategory{}, id)
	return result.Error
}
