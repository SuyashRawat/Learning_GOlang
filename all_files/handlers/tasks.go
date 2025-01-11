package handlers

import (
	"net/http"
	"taskmanage/models"

	"github.com/gin-gonic/gin"
)

type TaskInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Status      string `json:"status"`
	ProjectID   int    `json:"project_id" binding:"required"`
	AssignedTo  int    `json:"assigned_to"`
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       input.Title,
		Description: input.Description,
		Status:      input.Status,
		ProjectID:   input.ProjectID,
		AssignedTo:  input.AssignedTo,
	}

	if err := models.DB.Create(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}

// GetTasks fetches all tasks
func GetTasks(c *gin.Context) {
	var tasks []models.Task
	if err := models.DB.Find(&tasks).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tasks})
}

// GetTask fetches a task by ID
func GetTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": task})
}

// UpdateTask updates a task
func UpdateTask(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := models.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var input TaskInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task.Title = input.Title
	task.Description = input.Description
	task.Status = input.Status
	task.ProjectID = input.ProjectID
	task.AssignedTo = input.AssignedTo

	if err := models.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": task})
}
