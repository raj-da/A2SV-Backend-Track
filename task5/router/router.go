package router

import (
	"task-manager/controllers"

	"github.com/gin-gonic/gin"
)


func SetupRouter() *gin.Engine{
	router := gin.Default()
	router.GET("/task", controllers.GetTasks)
	router.GET("/task/:id", controllers.GetTask)
	router.PUT("/task/:id", controllers.UpdateTask)
	router.DELETE("/task/:id", controllers.DeleteTask)
	router.POST("/task", controllers.CreateTask)

	return router
}