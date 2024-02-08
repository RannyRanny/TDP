package repository

import (
	"encoding/json"
	"tdp-api/internal/model"

	"gorm.io/gorm"
)

type TemplateDayRepository interface {
	Create(templateDay model.TemplateDay) (model.TemplateDay, error)
	FindByID(id uint) (model.TemplateDay, error)
	Update(templateDay model.TemplateDay) (model.TemplateDay, error)
	Delete(id uint) error
}

type GormTemplateDayRepository struct {
	db *gorm.DB
}

func NewGormTemplateDayRepository(db *gorm.DB) TemplateDayRepository {
	return &GormTemplateDayRepository{db: db}
}

func (r *GormTemplateDayRepository) Create(templateDay model.TemplateDay) (model.TemplateDay, error) {

	categoriesJSON, err := json.Marshal(templateDay.TaskCategories)
	if err != nil {
		return templateDay, err
	}

	templateDay.TaskCategoriesJson = categoriesJSON
	result := r.db.Create(templateDay)
	return templateDay, result.Error
}

func (r *GormTemplateDayRepository) FindByID(id uint) (model.TemplateDay, error) {
	var templateDay model.TemplateDay
	result := r.db.Preload("TaskCategories").First(&templateDay, id)
	return templateDay, result.Error
}

func (r *GormTemplateDayRepository) FindByUserID(id uint) (model.TemplateDay, error) {
	var templateDay model.TemplateDay
	result := r.db.Preload("TemplateDay").First(&templateDay, id)
	return templateDay, result.Error
}

func (r *GormTemplateDayRepository) Update(templateDay model.TemplateDay) (model.TemplateDay, error) {
	result := r.db.Save(&templateDay)
	return templateDay, result.Error
}

func (r *GormTemplateDayRepository) Delete(id uint) error {
	result := r.db.Delete(&model.TemplateDay{}, id)
	return result.Error
}
