package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"
)

// DayHandler структура для обработчиков задач
type DayHandler struct {
	repo repository.IDayRepository[model.Day]
}

// NewDayHandler создает экземпляр DayHandler
func NewDayHandler(repo repository.IDayRepository[model.Day]) *DayHandler {
	return &DayHandler{repo: repo}
}

// CreateDay godoc
// @Summary Create a new template day
// @Description Add a new template day with categories
// @Tags Days
// @Accept json
// @Produce json
// @Param Day body model.Day true "Create Template Day"
// @Success 201 {object} model.Day
// @Failure 400 {object} object  "Not found"
// @Failure 500 {object} object "internal server error"
// @Router /day [post]
func (h *DayHandler) CreateDay(c *gin.Context) {
	var Day model.Day
	if err := c.ShouldBindJSON(&Day); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdDay, err := h.repo.Create(Day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdDay)
}

// GetDay godoc
// @Summary Get a template day
// @Description Get details of a template day by ID
// @Tags Days
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Success 200 {object} model.Day
// @Failure 404 {object} object  "Not found"
// @Router /day/{id} [get]
func (h *DayHandler) GetDay(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	Day, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "template day not found"})
		return
	}

	c.JSON(http.StatusOK, Day)
}

// UpdateDay godoc
// @Summary Update a template day
// @Description Update a template day by ID
// @Tags Days
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Param Day body model.Day true "Update Template Day"
// @Success 200 {object} model.Day
// @Failure 404 {object} object  "Not found"
// @Router /day/{id} [put]
func (h *DayHandler) UpdateDay(c *gin.Context) {
	var Day model.Day
	if err := c.ShouldBindJSON(&Day); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedDay, err := h.repo.Update(Day)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedDay)
}

// DeleteDay godoc
// @Summary Delete a template day
// @Description Delete a template day by ID
// @Tags Days
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Success 204 "No Content"
// @Failure 404 {object} object  "Not found"
// @Router /day/{id} [delete]
func (h *DayHandler) DeleteDay(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = h.repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetDays godoc
// @Summary Get a template day
// @Description Get details of a template day by ID
// @Tags Days
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Success 200 {array} model.Day
// @Failure 404 {object} object  "Not found"
// @Router /day/user/{userId} [get]
func (h *DayHandler) GetDays(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Param("userId"))
	tasks, err := h.repo.GetDays(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)

}

// TemplateDayHandler структура для обработчиков задач
type TemplateDayHandler struct {
	repo repository.IDayRepository[model.TemplateDay]
}

// NewTemplateDayHandler создает экземпляр TemplateDayHandler
func NewTemplateDayHandler(repo repository.IDayRepository[model.TemplateDay]) *TemplateDayHandler {
	return &TemplateDayHandler{repo: repo}
}

// CreateTemplateDay godoc
// @Summary Create a new template day
// @Description Add a new template day with categories
// @Tags templateDays
// @Accept json
// @Produce json
// @Param templateDay body model.TemplateDay true "Create Template Day"
// @Success 201 {object} model.TemplateDay
// @Failure 400 {object} object  "Not found"
// @Failure 500 {object} object "internal server error"
// @Router /day/template [post]
func (h *TemplateDayHandler) CreateTemplateDay(c *gin.Context) {
	var templateDay model.TemplateDay
	if err := c.ShouldBindJSON(&templateDay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdTemplateDay, err := h.repo.Create(templateDay)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdTemplateDay)
}

// GetTemplateDay godoc
// @Summary Get a template day
// @Description Get details of a template day by ID
// @Tags templateDays
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Success 200 {object} model.TemplateDay
// @Failure 404 {object} object  "Not found"
// @Router /day/template/{id} [get]
func (h *TemplateDayHandler) GetTemplateDay(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	templateDay, err := h.repo.FindByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "template day not found"})
		return
	}

	c.JSON(http.StatusOK, templateDay)
}

// UpdateTemplateDay godoc
// @Summary Update a template day
// @Description Update a template day by ID
// @Tags templateDays
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Param templateDay body model.TemplateDay true "Update Template Day"
// @Success 200 {object} model.TemplateDay
// @Failure 404 {object} object  "Not found"
// @Router /day/template/{id} [put]
func (h *TemplateDayHandler) UpdateTemplateDay(c *gin.Context) {
	var templateDay model.TemplateDay
	if err := c.ShouldBindJSON(&templateDay); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedTemplateDay, err := h.repo.Update(templateDay)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedTemplateDay)
}

// DeleteTemplateDay godoc
// @Summary Delete a template day
// @Description Delete a template day by ID
// @Tags templateDays
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Success 204 "No Content"
// @Failure 404 {object} object  "Not found"
// @Router /day/template/{id} [delete]
func (h *TemplateDayHandler) DeleteTemplateDay(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}

	err = h.repo.Delete(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

// GetTemplateDays godoc
// @Summary Get a template day
// @Description Get details of a template day by ID
// @Tags templateDays
// @Accept json
// @Produce json
// @Param id path int true "Template Day ID"
// @Success 200 {array} model.TemplateDay
// @Failure 404 {object} object  "Not found"
// @Router /day/template/user/{userId} [get]
func (h *TemplateDayHandler) GetTemplateDays(c *gin.Context) {

	userId, _ := strconv.Atoi(c.Param("userId"))
	tasks, err := h.repo.GetDays(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}
