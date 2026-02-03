package router

import (
	"task-manager/data"

	"github.com/gin-gonic/gin"
)


func InitializeRouter() {
	router := gin.Default()
	router.GET("/task", data.GetAllTasks)
	router.GET("/task/:id", data.GetTaskByID)
	router.PUT("/task/:id", data.UpdateTask)
	router.DELETE("/task/:id", data.DeleteTask)
	router.POST("/task", data.CreateTask)

	router.Run()
}