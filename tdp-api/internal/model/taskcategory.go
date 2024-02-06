package model

import (
	"time"
)

// TaskCategory структура, представляющая категорию задачи
type TaskCategory struct {
	ID          uint       `json:"id" gorm:"primarykey"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at" gorm:"index"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
}
