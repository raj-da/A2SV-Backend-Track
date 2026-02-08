package router

import (
	"task-manager/controllers"
	"task-manager/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Public Routes
	router.POST("/register", controllers.Register)
	router.POST("/login", controllers.Login)

	// Protected Routes (All Users)
	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.GET("/tasks", controllers.GetTasks)
		auth.GET("/tasks/:id", controllers.GetTask)

		// Admin only Routes
		admin := auth.Group("/")
		admin.Use(middleware.AdminMiddleware())
		{
			admin.POST("/tasks", controllers.CreateTask)
			admin.PUT("/tasks/:id", controllers.UpdateTask)
			admin.DELETE("/tasks/:id", controllers.DeleteTask)
			admin.PATCH("/promote/:username", controllers.Promote)
		}
	}

	return router
}