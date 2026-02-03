package controllers

import (
	"net/http"
	"task-manager/data"
	"task-manager/models"

	"github.com/gin-gonic/gin"
)

// GetAllTasks returns json representaton of all tasks
func GetTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"task": data.GetAllTasks()})
}

// GetTaskByID returns a json representation of a task
func GetTask(c *gin.Context) {
	id := c.Param("id") // get id from the path
	task, err := data.GetTaskByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask updates a task
func UpdateTask(c *gin.Context) {
	id := c.Param("id") // get ID from the path

	// convert json to Task struct
	var updatedTask models.Task
	if err := c.ShouldBindJSON(&updatedTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := data.UpdateTask(id, updatedTask); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task update successfully"})
}

// DeleteTask removes a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	if err := data.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	data.CreateTask(newTask)
	c.JSON(http.StatusCreated, newTask)
}