package handler

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sort"
	"strings"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"
)

// TelegramHandler структура для обработчиков задач
type TelegramHandler struct {
	repo repository.TelegramAuthRepository
}

// NewTelegramHandler создает экземпляр TelegramHandler
func NewTelegramHandler(repo repository.TelegramAuthRepository) *TelegramHandler {
	return &TelegramHandler{repo: repo}
}

// TelegramAuth godoc
// @Summary Telegram Authentication
// @Description Authenticates a user through Telegram data and saves or updates the user's data in the repository.
// @Tags telegram
// @Accept  json
// @Produce  json
// @Param data body model.TelegramAuthData true "Data for authentication"
// @Success 200 {integer} int64 "The Telegram User ID of the authenticated user"
// @Failure 400 {object} object "description of the error"
// @Failure 500 {object} object "Invalid data signature or internal server error"
// @Router /telegram/auth [post]
func (h *TelegramHandler) TelegramAuth(c *gin.Context) {
	var data model.TelegramAuthData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	botToken := ""

	if checkTelegramAuthData(data, botToken) {
		_, err := h.repo.FindByUserID(data.ID)
		if err != nil {
			h.repo.Save(data)
		}
		c.JSON(http.StatusOK, data.ID)
		return
	}
	c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid data signature"})
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
