package handler

import (
	"net/http"
	"strconv"
	"tdp-api/internal/model"
	"tdp-api/internal/repository"

	"github.com/gin-gonic/gin"
)

// TaskHandler структура для обработчиков задач
type TaskHandler struct {
	repo repository.TaskRepository
}

// NewTaskHandler создает экземпляр TaskHandler
func NewTaskHandler(repo repository.TaskRepository) *TaskHandler {
	return &TaskHandler{repo: repo}
}

// GetTasks обрабатывает GET запросы для получения списка задач
// @Summary Get list of tasks
// @Description Get details of all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Task
// @Router /tasks [get]
func (h *TaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.repo.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask обрабатывает GET запросы для получения одной задачи
// @Summary Get a task
// @Description Get details of a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} model.Task
// @Failure 404 {object} object "Not found"
// @Router /tasks/{id} [get]
func (h *TaskHandler) GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.repo.GetTask(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask обрабатывает POST запросы для создания новой задачи
// CreateTask godoc
// @Summary Create a new task
// @Description add by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body model.Task true "Add Task"
// @Success 200 {object} model.Task
// @Router /tasks [post]
func (h *TaskHandler) CreateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := h.repo.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask обрабатывает PUT запросы для обновления задачи
// @Summary Update a task
// @Description Update task details by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param task body model.Task true "Task body to update"
// @Success 200 {object} model.Task
// @Failure 400 {object} object "Bad Request"
// @Failure 404 {object} object "Not found"
// @Router /tasks/{id} [put]
func (h *TaskHandler) UpdateTask(c *gin.Context) {
	var task model.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	task.ID = uint(id)
	updatedTask, err := h.repo.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask обрабатывает DELETE запросы для удаления задачи
// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 204 "No Content"
// @Failure 404 {object} object "Not found"
// @Router /tasks/{id} [delete]
func (h *TaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}

// TemplateTaskHandler структура для обработчиков задач
type TemplateTaskHandler struct {
	repo repository.TemplateTaskRepository
}

// NewTemplateTaskHandler создает экземпляр TemplateTaskHandler
func NewTemplateTaskHandler(repo repository.TemplateTaskRepository) *TemplateTaskHandler {
	return &TemplateTaskHandler{repo: repo}
}

// GetTasks обрабатывает GET запросы для получения списка задач
// @Summary Get list of tasks
// @Description Get details of all tasks
// @Tags tasks
// @Accept  json
// @Produce  json
// @Success 200 {array} model.Task
// @Router /tasks/template [get]
func (h *TemplateTaskHandler) GetTasks(c *gin.Context) {
	tasks, err := h.repo.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask обрабатывает GET запросы для получения одной задачи
// @Summary Get a task
// @Description Get details of a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 200 {object} model.TemplateTask
// @Failure 404 {object} object "Not found"
// @Router /tasks/template/{id} [get]
func (h *TemplateTaskHandler) GetTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	task, err := h.repo.GetTask(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, task)
}

// CreateTask обрабатывает POST запросы для создания новой задачи
// CreateTask godoc
// @Summary Create a new task
// @Description add by json task
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param task body model.TemplateTask true "Add Task"
// @Success 200 {object} model.TemplateTask
// @Router /tasks/template [post]
func (h *TemplateTaskHandler) CreateTask(c *gin.Context) {
	var task model.TemplateTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdTask, err := h.repo.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// UpdateTask обрабатывает PUT запросы для обновления задачи
// @Summary Update a task
// @Description Update task details by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Param task body model.TemplateTask true "Task body to update"
// @Success 200 {object} model.TemplateTask
// @Failure 400 {object} object "Bad Request"
// @Failure 404 {object} object "Not found"
// @Router /tasks/template/{id} [put]
func (h *TemplateTaskHandler) UpdateTask(c *gin.Context) {
	var task model.TemplateTask
	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	task.ID = uint(id)
	updatedTask, err := h.repo.UpdateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask обрабатывает DELETE запросы для удаления задачи
// DeleteTask godoc
// @Summary Delete a task
// @Description Delete a task by ID
// @Tags tasks
// @Accept  json
// @Produce  json
// @Param id path int true "Task ID"
// @Success 204 "No Content"
// @Failure 404 {object} object "Not found"
// @Router /tasks/template/{id} [delete]
func (h *TemplateTaskHandler) DeleteTask(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := h.repo.DeleteTask(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusNoContent)
}
