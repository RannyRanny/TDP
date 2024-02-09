package repository

import (
	"gorm.io/gorm"
)

// IDayRepository generic interface for repository operations
type IDayRepository[T any] interface {
	GetDays(userId uint) ([]T, error)
	Create(entity T) (T, error)
	FindByID(id uint) (T, error)
	Update(entity T) (T, error)
	Delete(id uint) error
}

type DayRepository[T any] struct {
	db *gorm.DB
}

func NewDayRepository[T any](db *gorm.DB) *DayRepository[T] {
	return &DayRepository[T]{db: db}
}

func (r *DayRepository[T]) GetDays(userId uint) ([]T, error) {
	var days []T
	result := r.db.Preload("TaskCategory").Where("user_id = ?", userId).Find(&days)
	return days, result.Error
}
func (r *DayRepository[T]) Create(day T) (T, error) {
	result := r.db.Create(&day)
	return day, result.Error
}

func (r *DayRepository[T]) FindByID(id uint) (T, error) {
	var day T
	result := r.db.Preload("User").First(&day, id)
	return day, result.Error
}

func (r *DayRepository[T]) Update(day T) (T, error) {
	result := r.db.Save(&day)
	return day, result.Error
}

func (r *DayRepository[T]) Delete(id uint) error {
	var day T
	_ = r.db.Preload("User").First(&day, id)
	result := r.db.Delete(&day, id)
	return result.Error
}
