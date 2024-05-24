package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"tdp-api/internal/model"
	service "tdp-api/internal/services"
)

// TelegramHandler структура для обработчиков задач
type TelegramHandler struct {
	service service.TelegramService
}

// NewTelegramHandler создает экземпляр TelegramHandler
func NewTelegramHandler(service service.TelegramService) *TelegramHandler {
	return &TelegramHandler{service: service}
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

	result, err := h.service.Authenticate(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, result)

	c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid data signature"})
}

// CreateUser godoc
// @Summary Telegram User Create
// @Tags telegram
// @Accept  json
// @Produce  json
// @Param data body model.TelegramAuthData true "Data for authentication"
// @Success 200 {integer} int64 "The Telegram User ID of the authenticated user"
// @Failure 400 {object} object "description of the error"
// @Failure 500 {object} object "Invalid data signature or internal server error"
// @Router /user/create [post]
func (h *TelegramHandler) CreateUser(c *gin.Context) {
	var data model.TelegramAuthData

	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.service.CreateUser(data)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, result)

	c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid data signature"})
}
