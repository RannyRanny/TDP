package model

import (
	"time"
)

type TemplateDay struct {
	ID                 uint       `json:"id" gorm:"primarykey"`
	CreatedAt          time.Time  `json:"created_at"`
	UpdatedAt          time.Time  `json:"updated_at"`
	DeletedAt          *time.Time `json:"deleted_at" gorm:"index"`
	StartHour          int
	EndHour            int
	TaskCategoriesJson []byte
	TaskCategories     map[int]TaskCategory `gorm:"-"`
	UserID             uint
	User               TelegramAuthData `json:"user" gorm:"foreignKey:UserID;references:ID"`
}
