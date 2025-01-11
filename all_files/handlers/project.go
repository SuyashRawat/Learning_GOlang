package handlers

import (
	// "errors"
	"net/http"
	"taskmanage/models"

	"github.com/gin-gonic/gin"
)

type ProjectInput struct {
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
	OwnerID     int    `json:"owner_id" binding:"required"`
}

// CreateProject creates a new project
func CreateProject(c *gin.Context) {
	var input ProjectInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	project := models.Project{
		Name:        input.Name,
		Description: input.Description,
		OwnerID:     input.OwnerID,
	}

	if err := models.DB.Create(&project).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create project"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": project})
}

// GetProjects fetches all projects
func GetProjects(c *gin.Context) {
	var projects []models.Project
	if err := models.DB.Find(&projects).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch projects"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": projects})
}

// GetProject fetches a project by ID
func GetProject(c *gin.Context) {
	id := c.Param("id")
	var project models.Project
	if err := models.DB.First(&project, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Project not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": project})
}
