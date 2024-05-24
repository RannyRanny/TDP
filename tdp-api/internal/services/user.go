package service

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"sort"
	"strings"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"
)

// TelegramService определяет интерфейс для бизнес-логики работы с Telegram.
type TelegramService interface {
	Authenticate(data model.TelegramAuthData) (string, error)
	CreateUser(data model.TelegramAuthData) (int64, error)
}

// TelegramServiceImpl структура сервиса для Telegram.
type TelegramServiceImpl struct {
	repo repository.TelegramAuthRepository
}

// NewTelegramService создает новый экземпляр сервиса Telegram.
func NewTelegramService(repo repository.TelegramAuthRepository) TelegramService {
	return &TelegramServiceImpl{repo: repo}
}

func (s *TelegramServiceImpl) Authenticate(data model.TelegramAuthData) (string, error) {

	botToken := ""

	if checkTelegramAuthData(data, botToken) {
		return botToken, nil
	}
	return "", errors.New("auth fail")
}

func (s *TelegramServiceImpl) CreateUser(data model.TelegramAuthData) (int64, error) {
	botToken := ""

	if checkTelegramAuthData(data, botToken) {
		user, err := s.repo.FindByUserID(data.ID)
		if err != nil {
			user, err = s.repo.Save(data)
			if err != nil {
				return int64(user.ID), nil
			}
		}
		
		return int64(user.ID), nil
	}
	return -1, errors.New("auth fail")
}

// Возвращает true, если подпись верна.
func checkTelegramAuthData(data model.TelegramAuthData, botToken string) bool {
	// Собираем данные для проверки подписи
	var keyValuePairs []string
	keyValuePairs = append(keyValuePairs, fmt.Sprintf("auth_date=%d", data.AuthDate))
	keyValuePairs = append(keyValuePairs, fmt.Sprintf("first_name=%s", data.FirstName))
	keyValuePairs = append(keyValuePairs, fmt.Sprintf("id=%d", data.ID))
	keyValuePairs = append(keyValuePairs, fmt.Sprintf("last_name=%s", data.LastName))
	keyValuePairs = append(keyValuePairs, fmt.Sprintf("photo_url=%s", data.PhotoURL))
	keyValuePairs = append(keyValuePairs, fmt.Sprintf("username=%s", data.Username))

	sort.Strings(keyValuePairs)
	dataCheckString := strings.Join(keyValuePairs, "\n")

	// Вычисляем HMAC-SHA256 хеш
	hmacInstance := hmac.New(sha256.New, []byte(botToken))
	hmacInstance.Write([]byte(dataCheckString))
	hash := hex.EncodeToString(hmacInstance.Sum(nil))

	// Сравниваем полученный хеш с тем, что пришел от Telegram
	return hash == data.Hash
}
