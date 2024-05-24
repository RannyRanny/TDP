package model

import (
	"encoding/json"
	"time"
)

// BaseDay структура, содержащая общие поля для TemplateDay и других подобных структур.
type BaseDay struct {
	ID        uint             `json:"id" gorm:"primarykey"`
	CreatedAt time.Time        `json:"created_at"`
	UpdatedAt time.Time        `json:"updated_at"`
	DeletedAt *time.Time       `json:"deleted_at" gorm:"index"`
	Date      time.Time        `json:"date"`
	StartHour int              `json:"start_hour"`
	EndHour   int              `json:"end_hour"`
	UserID    uint             `json:"user_id"`
	User      TelegramAuthData `json:"user" gorm:"foreignKey:UserID;references:ID"`
}

// TemplateDay расширяет BaseDay с полями, специфичными для шаблона дня.
type TemplateDay struct {
	BaseDay
	TaskCategoriesJson json.RawMessage      `json:"-" gorm:"type:json"`
	TaskCategories     map[int]TaskCategory `json:"task_categories" gorm:"-"`
	TasksJson          json.RawMessage      `json:"-" gorm:"type:json"`
	Tasks              map[int]TemplateTask `json:"tasks" gorm:"-"`
}

// Day расширяет BaseDay с полями, специфичными для шаблона дня.
type Day struct {
	BaseDay
	TaskCategoriesJson json.RawMessage      `json:"-" gorm:"type:json"`
	TaskCategories     map[int]TaskCategory `json:"task_categories" gorm:"-"`
	TasksJson          json.RawMessage      `json:"-" gorm:"type:json"`
	Tasks              map[int]Task         `json:"tasks" gorm:"-"`
}
