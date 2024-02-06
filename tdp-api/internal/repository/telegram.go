package repository

import (
	"gorm.io/gorm"
	"tdp-api/internal/model"
)

type TelegramAuthRepository interface {
	Save(authData model.TelegramAuthData) (model.TelegramAuthData, error)
	FindByUserID(userID int64) (model.TelegramAuthData, error)
}

type GormTelegramAuthRepository struct {
	db *gorm.DB
}

func NewGormTelegramAuthRepository(db *gorm.DB) TelegramAuthRepository {
	return &GormTelegramAuthRepository{db: db}
}

func (r *GormTelegramAuthRepository) Save(authData model.TelegramAuthData) (model.TelegramAuthData, error) {
	result := r.db.Save(&authData)
	return authData, result.Error
}

func (r *GormTelegramAuthRepository) FindByUserID(userID int64) (model.TelegramAuthData, error) {
	var authData model.TelegramAuthData
	result := r.db.Where("user_id = ?", userID).First(&authData)
	return authData, result.Error
}
