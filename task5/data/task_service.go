package data

import (
	"net/http"
	"task-manager/models"
	"time"

	"github.com/gin-gonic/gin"
)

// Dummy task data
var tasks = []models.Task{
    {ID: "1", Title: "Task 1", Description: "First task", DueDate: time.Now(), Status: "Pending"},
    {ID: "2", Title: "Task 2", Description: "Second task", DueDate: time.Now().AddDate(0, 0, 1), Status: "In Progress"},
    {ID: "3", Title: "Task 3", Description: "Third task", DueDate: time.Now().AddDate(0, 0, 2), Status: "Completed"},
}

// GetAllTasks returns json representaton of all tasks
func GetAllTasks(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"task": tasks})
}

// GetTaskByID returns a json representation of a task
func GetTaskByID(c *gin.Context) {
	id := c.Param("id") // get id from the path
	// iterate over the tasks and find the task
	for _, task := range tasks {
		if task.ID == id {
			c.JSON(http.StatusOK, task)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
}

// UpdateTask updates a task
func UpdateTask(c *gin.Context) {
	id := c.Param("id") // get ID from the path

	// convert json to Task struct
	var updatedTask models.Task
	if err := c.BindJSON(&updatedTask); err != nil {
		return
	}

	// Look for the specific task and update
	for ind, task := range tasks {
		if task.ID == id {
			if updatedTask.Title != "" {
				task.Title = updatedTask.Title
			}
			if updatedTask.Description != "" {
				task.Description = updatedTask.Description
			}
			tasks[ind] = task
			c.JSON(http.StatusOK, gin.H{"message": "Task updated"})
			return
		}
	}

	// If task not found
	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// DeleteTask removes a task
func DeleteTask(c *gin.Context) {
	id := c.Param("id")
	for ind, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:ind],tasks[ind+1:]...)
			c.JSON(http.StatusOK, gin.H{"message":"Task removed"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "Task not found"})
}

// CreateTask creates a new task
func CreateTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		return
	}
	tasks = append(tasks, newTask)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created"})
}