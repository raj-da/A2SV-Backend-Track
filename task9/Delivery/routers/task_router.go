package routers

import (
	"task-manager/Delivery/controllers"
	repository "task-manager/Repository"
	usecases "task-manager/Usecases"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

// NewTaskRouter sets up the dependencies and routes for Tasks
func NewTaskRouter(group *gin.RouterGroup, db *mongo.Database, collection string) {
	// 1. Create the data layer
	repo := repository.NewTaskRepository(
		db,
		collection,
	)

	// 2. Create usecase layer
	uc := usecases.NewTaskUsercase(repo)

	// 3. Create controller
	tc := &controllers.TaskController{TaskUsecase: uc}

	// 4. Map the paths
	group.POST("/tasks", tc.Create)
	group.GET("/tasks", tc.GetTasks)
	group.GET("/tasks/:id", tc.GetTask)
	group.DELETE("/tasks/:id", tc.Delete)
	group.PUT("/tasks/:id", tc.Update)
}