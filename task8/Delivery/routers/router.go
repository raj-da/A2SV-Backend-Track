package routers

import (
	infrastructure "task-manager/Infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

func SetUp(db *mongo.Database) *gin.Engine {
	router := gin.Default()

	taskCollection := "task"
	userCollection := "user"

	// PUBLIC ROUTES
	public := router.Group("/auth")
	NewUserRouter(public, db, userCollection)

	protected := router.Group("/api")
	// Add middleware to protected routs
	protected.Use(infrastructure.AuthMiddleware(), infrastructure.AdminMiddleware())
	NewTaskRouter(protected, db, taskCollection)

	return router
}