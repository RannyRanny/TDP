package handler

import (
	"net/http"
	"strconv"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"

	"github.com/gin-gonic/gin"
)

// TaskCategoryHandler структура для обработчиков категорий задач
type TaskCategoryHandler struct {
	repo repository.TaskCategoryRepository
}

// NewTaskCategoryHandler создает экземпляр TaskCategoryHandler
func NewTaskCategoryHandler(repo repository.TaskCategoryRepository) *TaskCategoryHandler {
	return &TaskCategoryHandler{repo: repo}
}

// GetCategories обрабатывает GET запросы для получения списка категорий
// @Summary Get list of task categories
// @Description Get details of all task categories
// @Tags categories
// @Accept  json
// @Produce  json
// @Success 200 {array} model.TaskCategory
// @Router /categories/user/{userId} [get]
func (h *TaskCategoryHandler) GetCategories(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	categories, err := h.repo.GetCategories(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, categories)
}

// GetCategory обрабатывает GET запросы для получения одной категории
// @Summary Get a task category
// @Description Get details of a task category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 200 {object} model.TaskCategory
// @Failure 404 {object} object "Not found"
// @Router /categories/{id} [get]
func (h *TaskCategoryHandler) GetCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	category, err := h.repo.GetCategory(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, category)
}

// CreateCategory обрабатывает POST запросы для создания новой категории
// @Summary Create a new task category
// @Description Add a new task category
// @Tags categories
// @Accept  json
// @Produce  json
// @Param category body model.TaskCategory true "Task Category to add"
// @Success 201 {object} model.TaskCategory
// @Router /categories [post]
func (h *TaskCategoryHandler) CreateCategory(c *gin.Context) {
	var category model.TaskCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdCategory, err := h.repo.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdCategory)
}

// UpdateCategory обрабатывает PUT запросы для обновления категории
// @Summary Update a task category
// @Description Update task category details by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Param category body model.TaskCategory true "Task Category body to update"
// @Success 200 {object} model.TaskCategory
// @Failure 400 {object} object "Bad Request"
// @Failure 404 {object} object "Not found"
// @Router /categories/{id} [put]
func (h *TaskCategoryHandler) UpdateCategory(c *gin.Context) {
	var category model.TaskCategory
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	category.ID = uint(id)
	updatedCategory, err := h.repo.UpdateCategory(category)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedCategory)
}

// DeleteCategory обрабатывает DELETE запросы для удаления категории
// @Summary Delete a task category
// @Description Delete a task category by ID
// @Tags categories
// @Accept  json
// @Produce  json
// @Param id path int true "Category ID"
// @Success 204 "No Content"
// @Failure 404 {object} object "Not found"
// @Router /categories/{id} [delete]
func (h *TaskCategoryHandler) DeleteCategory(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.DeleteCategory(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
