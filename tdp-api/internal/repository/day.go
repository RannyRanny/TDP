package repository

import (
	"tdp-api/internal/model"
	"time"

	"gorm.io/gorm"
)

type DayRepository interface {
	GetDays(userId uint) ([]model.Day, error)
	Create(day model.Day) (model.Day, error)
	FindByID(id uint) (model.Day, error)
	FindByDate(date time.Time, userId uint) ([]model.Day, error)
	Update(day model.Day) (model.Day, error)
	Delete(id uint) error
}

type TemplateDayRepository interface {
	GetDays(userId uint) ([]model.TemplateDay, error)
	Create(templateDay model.TemplateDay) (model.TemplateDay, error)
	FindByID(id uint) (model.TemplateDay, error)
	Update(templateDay model.TemplateDay) (model.TemplateDay, error)
	Delete(id uint) error
}

type GormDayRepository struct {
	db *gorm.DB
}

func NewGormDayRepository(db *gorm.DB) DayRepository {
	return &GormDayRepository{db: db}
}

func (r *GormDayRepository) GetDays(userId uint) ([]model.Day, error) {
	var days []model.Day
	result := r.db.Preload("TaskCategory").Where("user_id = ?", userId).Find(&days)
	return days, result.Error
}

func (r *GormDayRepository) Create(day model.Day) (model.Day, error) {
	result := r.db.Create(&day)
	return day, result.Error
}

func (r *GormDayRepository) FindByID(id uint) (model.Day, error) {
	var day model.Day
	result := r.db.Preload("User").First(&day, id)
	return day, result.Error
}

func (r *GormDayRepository) FindByDate(date time.Time, userId uint) ([]model.Day, error) {
	var days []model.Day

	result := r.db.Preload("TaskCategory").
		Where("user_id = ? AND DATE(date) = ?", userId, date.Format("2006-01-02")).
		Find(&days)

	if result.Error != nil {
		return days, result.Error
	}

	return days, nil
}

func (r *GormDayRepository) Update(day model.Day) (model.Day, error) {
	result := r.db.Save(&day)
	return day, result.Error
}

func (r *GormDayRepository) Delete(id uint) error {
	result := r.db.Delete(&model.TemplateDay{}, id)
	return result.Error
}

type GormTemplateDayRepository struct {
	db *gorm.DB
}

func (r *GormTemplateDayRepository) GetDays(userId uint) ([]model.TemplateDay, error) {
	var days []model.TemplateDay
	result := r.db.Preload("TaskCategory").Where("user_id = ?", userId).Find(&days)
	return days, result.Error
}

func NewGormTemplateDayRepository(db *gorm.DB) TemplateDayRepository {
	return &GormTemplateDayRepository{db: db}
}

func (r *GormTemplateDayRepository) Create(templateDay model.TemplateDay) (model.TemplateDay, error) {
	result := r.db.Create(&templateDay)
	return templateDay, result.Error
}

func (r *GormTemplateDayRepository) FindByID(id uint) (model.TemplateDay, error) {
	var templateDay model.TemplateDay
	result := r.db.Preload("User").First(&templateDay, id)
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
