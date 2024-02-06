package model

import (
	"time"
)

// Task структура, представляющая сущность "Задача"
type Task struct {
	ID           uint         `json:"id" gorm:"primarykey"`
	CreatedAt    time.Time    `json:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	DeletedAt    *time.Time   `json:"deleted_at" gorm:"index"`
	Title        string       `json:"title"`
	Description  string       `json:"description"`
	IsCompleted  bool         `json:"is_completed"`
	CategoryID   uint         `json:"category_id"`
	TaskCategory TaskCategory `json:"category" gorm:"foreignKey:CategoryID;references:ID"` // Связь с TaskCategory
}
