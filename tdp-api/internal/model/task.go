package model

import (
	"time"
)

// BaseTask структура, представляющая базовые поля задачи
type BaseTask struct {
	ID           uint             `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time        `json:"created_at"`
	UpdatedAt    time.Time        `json:"updated_at"`
	DeletedAt    *time.Time       `json:"deleted_at" gorm:"index"`
	Title        string           `json:"title"`
	Description  string           `json:"description"`
	IsCompleted  bool             `json:"is_completed"`
	CategoryID   uint             `json:"category_id"`
	TaskCategory TaskCategory     `json:"category" gorm:"foreignKey:CategoryID;references:ID"`
	UserID       uint             `json:"user_id"`
	User         TelegramAuthData `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

// Task структура, представляющая сущность "Задача"
type Task struct {
	BaseTask // Встраивание BaseTask для наследования полей
}

// TemplateTask структура, представляющая сущность "Шаблон задачи"
type TemplateTask struct {
	BaseTask // Встраивание BaseTask для наследования полей
}
